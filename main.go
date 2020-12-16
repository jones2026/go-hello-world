package main

import (
	"encoding/json"
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

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/healthz", HealthHandler)

	port := ":8080"
	log.Println("Listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, r))
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
