package handler

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/models"
	"github.com/SANEKNAYMCHIK/mock-cbr/internal/services"
)

func GetRates(w http.ResponseWriter, r *http.Request) {
	testId := r.URL.Query().Get("test_id")
	status := r.URL.Query().Get("status")
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		// log.Printf("Incorrect status in request")
		http.Error(w, "Incorrect status in request", http.StatusInternalServerError)
		return
	}
	if statusCode != http.StatusOK {
		// log.Printf("unsuccessful data receiving")
		http.Error(w, "Unsuccessful data receiving", http.StatusInternalServerError)
		return
	}
	WriteResponse(w, services.GetRates(testId, statusCode))
	// func(w http.ResponseWriter, r *http.Request) {

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"USD": 1.0, "EUR": 0.85, "GBP": 0.75}`))
	// }
}

func WriteResponse(w http.ResponseWriter, rate models.Rate) {
	if err := xml.NewEncoder(w).Encode(rate); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
