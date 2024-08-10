package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Number1 *int `json:"number1"`
		Number2 *int `json:"number2"`
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

	if requestBody.Number1 == nil || requestBody.Number2 == nil {
		slog.Error("Both numbers are required")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Both number1 and number2 are required",
		})
		return
	}

	result := *requestBody.Number1 * *requestBody.Number2

	json.NewEncoder(w).Encode(map[string]int{
		"result": result,
	})
}