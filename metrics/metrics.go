package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	PromRequestStatusIncoming = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_status_incoming_total",
			Help: "Total requests where (service) responded (status)",
		},
		[]string{"service", "status"},
	)
)

func init() {
	log.Println("global metrics initializing")
	prometheus.MustRegister(PromRequestStatusIncoming)
}
