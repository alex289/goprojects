package handler

import (
	"calcapi/db"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Numbers []float64 `json:"numbers"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		slog.Error("Error decoding request body")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"result": "Invalid request body",
		})
		return
	}

	if requestBody.Numbers == nil {
		slog.Error("A list of numbers is required")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "A list of numbers is required",
		})
		return
	}

	var sum float64 = 0

	for _, number := range requestBody.Numbers {
		sum += number
	}

	db.TrackRequest(r, w, -1, -1, "sum", sum)

	json.NewEncoder(w).Encode(map[string]float64{
		"result": sum,
	})
}
