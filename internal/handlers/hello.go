package handlers

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
