package main

import "github.com/prometheus/client_golang/prometheus"

var (
	currentConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "sshforever_connections_current",
		Help: "The current number of connections.",
	})

	timeWasted = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "sshforever_time_wasted_milliseconds",
		Help:    "The time clients stayed connected in milliseconds.",
		Buckets: prometheus.LinearBuckets(60000, 60000, 10), // 10 buckets, each 60 seconds wide.
	})
)

func registerMetrics() {
	prometheus.MustRegister(currentConnections)
	prometheus.MustRegister(timeWasted)
}
