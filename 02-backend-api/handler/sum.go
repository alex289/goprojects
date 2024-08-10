package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Numbers []int `json:"numbers"`
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

	sum := 0

	for _, number := range requestBody.Numbers {
		sum += number
	}

	json.NewEncoder(w).Encode(map[string]int{
		"result": sum,
	})
}
