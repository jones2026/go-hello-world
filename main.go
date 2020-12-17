package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jones2026/go-hello-world/internal/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	slokmdlw "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

func main() {
	log.Println("Starting app...")
	r := chi.NewRouter()

	prometheusMiddleware := slokmdlw.New(slokmdlw.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r.Use(std.HandlerProvider("", prometheusMiddleware))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Hello)
	r.Get("/healthz", handlers.Health)
	r.Handle("/metrics", promhttp.Handler())

	port := ":8080"
	log.Println("Listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
