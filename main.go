package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jones2026/go-hello-world/internal/handlers"
)

func main() {
	log.Println("Starting app...")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Hello)
	r.Get("/healthz", handlers.Health)

	port := ":8080"
	log.Println("Listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, r))
}

