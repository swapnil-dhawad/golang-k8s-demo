package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	// Register custom metrics with the global Prometheus registry
	prometheus.MustRegister(httpRequestsTotal)
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	// Increment the counter for the welcome page
	httpRequestsTotal.With(prometheus.Labels{"path": "/"}).Inc()

	// Render the welcome html page
	http.ServeFile(w, r, "static/welcome.html")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the Prometheus metrics
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	// Serve the welcome page at the root URL
	http.HandleFunc("/", welcomePage)

	// Serve the Prometheus metrics at the /metrics URL
	http.HandleFunc("/metrics", metricsHandler)

	// Start the HTTP server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
