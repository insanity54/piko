package agent

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	// ForwardRequestsTotal is the total number of requests send to the
	// forward addr. Labelled by method, status code and endpoint ID.
	ForwardRequestsTotal *prometheus.CounterVec
	// ForwardErrorsTotal is the total number of errors from the forward
	// address (not including bad status codes). Labelled by endpoint ID.
	ForwardErrorsTotal *prometheus.CounterVec
	// ForwardRequestLatency is a histogram of the latency of requests to the
	// forward address in seconds. Labelled by response status code and
	// endpoint ID.
	ForwardRequestLatency *prometheus.HistogramVec

	// ServerConnectionErrorsTotal is the total number of errors attempting to
	// connect to the server.
	ServerConnectionErrorsTotal prometheus.Counter
}

func NewMetrics() *Metrics {
	return &Metrics{
		ForwardRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "forward",
				Name:      "requests_total",
				Help:      "Total requests to the forward address.",
			},
			[]string{"method", "status", "endpoint_id"},
		),
		ForwardErrorsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "forward",
				Name:      "errors_total",
				Help:      "Total errors from the forward address.",
			},
			[]string{"endpoint_id"},
		),
		ForwardRequestLatency: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "forward",
				Name:      "request_latency_seconds",
				Help:      "Forward request latency in seconds",
				Buckets:   prometheus.ExponentialBuckets(0.01, 2, 10),
			},
			[]string{"status", "endpoint_id"},
		),
	}
}

func (m *Metrics) Register(registry *prometheus.Registry) {
	registry.MustRegister(
		m.ForwardRequestsTotal,
		m.ForwardErrorsTotal,
		m.ForwardRequestLatency,
	)
}
