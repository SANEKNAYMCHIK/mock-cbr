package main

import (
	"log"
	"net/http"

	_ "github.com/SANEKNAYMCHIK/mock-cbr/docs"
	"github.com/SANEKNAYMCHIK/mock-cbr/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample mock server.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		handler.GetRates(w, r)
	})
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
