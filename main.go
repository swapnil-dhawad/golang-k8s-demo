package main

import (
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "handler"},
    )
)

func init() {
    prometheus.MustRegister(httpRequests)
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
    httpRequests.WithLabelValues(r.Method, "welcomePage").Inc()
    // Render the course html page
    http.ServeFile(w, r, "static/welcome.html")
}

func main() {
    http.HandleFunc("/", welcomePage)
    http.Handle("/metrics", promhttp.Handler()) // Expose metrics endpoint
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
