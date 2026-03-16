package main

import (
	"net/http"
	"sync"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		handler.GetRates(w, r)
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func(mux *http.ServeMux) {
		defer wg.Done()
		http.ListenAndServe(":8080", mux)
	}(mux)
	wg.Wait()
}
