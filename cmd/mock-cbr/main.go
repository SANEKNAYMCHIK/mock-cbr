package main

import (
	"log"
	"net/http"

	_ "github.com/SANEKNAYMCHIK/mock-cbr/docs"
	"github.com/SANEKNAYMCHIK/mock-cbr/internal/handler"
	"github.com/SANEKNAYMCHIK/mock-cbr/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample mock server.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rate", handler.GetRates)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	loggedMux := handler.LoggingMiddleware(mux)
	service.InitRedis()
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
