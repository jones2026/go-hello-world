package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type healthResponse struct {
	Hostname string
	Metadata map[string]string
}

//Health returns 200 status code
func Health(w http.ResponseWriter, r *http.Request) {
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