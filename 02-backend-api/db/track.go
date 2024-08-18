package db

import (
	"calcapi/middleware"
	"net/http"
)

func TrackRequest(
	r *http.Request,
	w http.ResponseWriter,
	number1 float64,
	number2 float64,
	operator string,
	result float64) error {
	requestID, ok := middleware.FromContext(r.Context())

	if !ok {
		http.Error(w, "Request ID not found", http.StatusInternalServerError)
	}

	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `
        INSERT INTO calculations (number1, number2, operator, result, request_id)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err = db.Exec(query, number1, number2, operator, result, requestID)
	return err
}
