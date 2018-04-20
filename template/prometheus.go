package main

// TODO: create init prometheus stuff for monitoring
import (
	"net/http"

	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

type PrometheusMetricsController struct {
	port int
	addr string
}

func newPrometheusMetricsController(port int) *PrometheusMetricsController {
	p := &PrometheusMetricsController{
		port: port,
		addr: fmt.Sprintf(":%v", port),
	}
	return p
}

func (p *PrometheusMetricsController) registerMetrics() {
	prometheus.MustRegister(metricActiveServicesEventsCounter)
}

func (p *PrometheusMetricsController) Run(stopChan chan struct{}) {
	log().Info("Starting PrometheusMetricsController")

	p.registerMetrics()
	p.httpServer()

	<-stopChan
	log().Info("Stopping PrometheusMetricsController")
}

func (p *PrometheusMetricsController) httpServer() {
	http.Handle("/metrics", promhttp.Handler())
	log().Fatal(http.ListenAndServe(p.addr, nil))
}
