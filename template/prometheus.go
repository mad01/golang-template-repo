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

type PrometheusController struct {
	port int
	addr string
}

func newPrometheusController(port int) *PrometheusController {
	p := &PrometheusController{
		port: port,
		addr: fmt.Sprintf(":%v", port),
	}
	return p
}

func (p *PrometheusController) registerMetrics() {
	prometheus.MustRegister(metricActiveServicesEventsCounter)
}

func (p *PrometheusController) Run(stopChan chan struct{}) {
	log().Info("Starting PrometheusController")

	p.registerMetrics()
	p.httpServer()

	<-stopChan
	log().Info("Stopping PrometheusController")
}

func (p *PrometheusController) httpServer() {
	http.Handle("/metrics", promhttp.Handler())
	log().Fatal(http.ListenAndServe(p.addr, nil))
}

func (p *PrometheusController) handler() http.Handler {
	return promhttp.Handler()
}
