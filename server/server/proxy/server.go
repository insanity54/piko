package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/andydunstall/pico/pkg/log"
	"github.com/andydunstall/pico/pkg/status"
	"github.com/andydunstall/pico/server/config"
	"github.com/andydunstall/pico/server/proxy"
	"github.com/andydunstall/pico/server/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

// Server is the HTTP server for the proxy, which proxies all incoming
// requests.
type Server struct {
	addr string

	router *gin.Engine

	httpServer *http.Server

	proxy *proxy.Proxy

	shutdownCtx    context.Context
	shutdownCancel func()

	conf *config.ProxyConfig

	logger log.Logger
}

func NewServer(
	addr string,
	proxy *proxy.Proxy,
	conf *config.ProxyConfig,
	registry *prometheus.Registry,
	logger log.Logger,
) *Server {
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())

	router := gin.New()
	server := &Server{
		addr:   addr,
		router: router,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		shutdownCtx:    shutdownCtx,
		shutdownCancel: shutdownCancel,
		proxy:          proxy,
		conf:           conf,
		logger:         logger.WithSubsystem("proxy.server"),
	}

	// Recover from panics.
	server.router.Use(gin.CustomRecovery(server.panicRoute))

	server.router.Use(middleware.NewLogger(logger))
	if registry != nil {
		router.Use(middleware.NewMetrics("proxy", registry))
	}

	server.registerRoutes()

	return server
}

func (s *Server) Serve() error {
	s.logger.Info("starting http server", zap.String("addr", s.addr))

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http serve: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) registerRoutes() {
	// Handle not found routes, which includes all proxied endpoints.
	s.router.NoRoute(s.notFoundRoute)
}

// proxyRoute handles proxied requests from downstream clients.
func (s *Server) proxyRoute(c *gin.Context) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(s.conf.GatewayTimeout)*time.Second,
	)
	defer cancel()

	resp, err := s.proxy.Request(ctx, c.Request)
	if err != nil {
		var errorInfo *status.ErrorInfo
		if errors.As(err, &errorInfo) {
			c.JSON(errorInfo.StatusCode, gin.H{"error": errorInfo.Message})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	// Write the response status, headers and body.
	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		s.logger.Warn("failed to write response", zap.Error(err))
	}
}

func (s *Server) notFoundRoute(c *gin.Context) {
	// All /pico endpoints are reserved. All others are proxied.
	if strings.HasPrefix(c.Request.URL.Path, "/pico") {
		c.Status(http.StatusNotFound)
		return
	}
	s.proxyRoute(c)
}

func (s *Server) panicRoute(c *gin.Context, err any) {
	s.logger.Error(
		"handler panic",
		zap.String("path", c.FullPath()),
		zap.Any("err", err),
	)
	c.AbortWithStatus(http.StatusInternalServerError)
}

func init() {
	// Disable Gin debug logs.
	gin.SetMode(gin.ReleaseMode)
}
