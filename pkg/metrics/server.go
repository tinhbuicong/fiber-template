// Package metrics exposes Prometheus /metrics endpoint (App → Prometheus → Alertmanager).
package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Serve starts an HTTP server on addr (e.g. "0.0.0.0:9091") serving only GET /metrics
// for Prometheus scrape. Chạy trong goroutine từ main.
func Serve(addr string) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{Addr: addr, Handler: mux}
	log.Printf("[metrics] listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("[metrics] server error: %v", err)
	}
}
