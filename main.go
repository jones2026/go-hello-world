package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/jones2026/go-hello-world/healthz"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type hello struct {
	Name string
	Time string
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	log.Println("Starting app...")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	
	versionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		fmt.Fprintf(w, "1.0")
	})

	defaultHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UTC()
		stamp := ("The current machine timestamp in UTC: " + t.Format("2006-01-02T15:04:05.999999-07:00"))
		fmt.Println(stamp)
		fmt.Printf("Received %v request for %v\n", r.Method, r.URL)
		fmt.Fprintf(w, stamp)
	})

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hello := hello{"Anonymous", time.Now().Format(time.Stamp)}
		templates := template.Must(template.ParseFiles("./templates/hello-template.html"))
		if name := r.FormValue("name"); name != "" {
			hello.Name = name
		}
		fmt.Printf("Received %v request for %v from: %v\n", r.Method, r.URL, hello)
		if err := templates.ExecuteTemplate(w, "hello-template.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte("Pull"))
	})

	hc := &healthz.Config{
		Hostname: hostname,
	}
	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	recordMetrics()

	http.Handle("/stylesheets/", prometheus.InstrumentHandler(
		"stylesheets", http.FileServer(http.Dir("./static"))))

	http.Handle("/", prometheus.InstrumentHandler("default", defaultHandler))
	http.Handle("/version", prometheus.InstrumentHandler("version", versionHandler))
	http.Handle("/hello", prometheus.InstrumentHandler("hello", helloHandler))
	http.Handle("/healthz", prometheus.InstrumentHandler("healthz", healthzHandler))
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":80", nil))
}
