package handler

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/models"
	"github.com/SANEKNAYMCHIK/mock-cbr/internal/service"
)

func GetRates(w http.ResponseWriter, r *http.Request) {
	testId := r.URL.Query().Get("test_id")
	status := r.URL.Query().Get("status")
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		http.Error(w, "Incorrect status in request", http.StatusInternalServerError)
		return
	}
	if statusCode != http.StatusOK {
		http.Error(w, "Unsuccessful data receiving", http.StatusInternalServerError)
		return
	}
	WriteResponse(w, service.GetRates(testId, statusCode))
}

func WriteResponse(w http.ResponseWriter, rate models.Rate) {
	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(rate); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
