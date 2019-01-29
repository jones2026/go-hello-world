package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/jones2026/go-hello-world/healthz"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type hello struct {
	Name string
	Time string
}

type config struct {
}

func main() {
	log.Println("Starting app...")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UTC()
		stamp := ("The current machine timestamp in UTC: " + t.Format("2006-01-02T15:04:05.999999-07:00"))
		// stamp
		fmt.Println(stamp)
		fmt.Printf("Received %v request for %v\n", r.Method, r.URL)
		fmt.Fprintf(w, stamp)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello := hello{"Anonymous", time.Now().Format(time.Stamp)}
		templates := template.Must(template.ParseFiles("./templates/hello-template.html"))
		if name := r.FormValue("name"); name != "" {
			hello.Name = name
		}
		fmt.Printf("Received %v request for %v from: %v\n", r.Method, r.URL, hello)
		if err := templates.ExecuteTemplate(w, "hello-template.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	hc := &healthz.Config{
		Hostname: hostname,
	}
	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/healthz", healthzHandler)

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":80", nil))
}
