package main

import (
	"log"
	"net/http"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		handler.GetRates(w, r)
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
