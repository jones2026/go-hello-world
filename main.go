package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type healthResponse struct {
	Hostname string
	Metadata map[string]string
}

func main() {
	log.Println("Starting app...")

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/healthz", HealthHandler)

	fmt.Println("Listening")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

//HealthHandler returns 200 status code
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	metadata := make(map[string]string)
	metadata["timestamp"] = time.Now().Format(time.Stamp)
	metadata["go_version"] = runtime.Version()
	metadata["os_type"] = runtime.GOOS
	metadata["arch"] = runtime.GOARCH

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err.Error())
	}

	response := healthResponse{
		Hostname: hostname,
		Metadata: metadata,
	}
	data, err := json.MarshalIndent(&response, "", "  ")
	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Write(data)
}
