package main

// TODO: create init prometheus stuff for monitoring
import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix = "{{GITHUB_REPO}}"
)

var (
	// active services
	metricActiveServicesEventsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: getMetricPrefix("active_services_events"),
		Help: "number of service events"},
		[]string{},
	)
)

func getMetricPrefix(name string) string {
	return fmt.Sprintf("%v_%v", metricPrefix, name)
}

func init() {
	prometheus.MustRegister(metricPatchedPodsCounter)
}
