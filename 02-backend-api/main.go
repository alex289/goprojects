package main

import (
	"calcapi/handler"
	"calcapi/middleware"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /add", handler.AddHandler)
	router.HandleFunc("POST /subtract", handler.SubtractHandler)
	router.HandleFunc("POST /multiply", handler.MultiplyHandler)
	router.HandleFunc("POST /divide", handler.DivideHandler)
	router.HandleFunc("POST /sum", handler.SumHandler)

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Headers,
	)

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(cors.Default().Handler(router)),
	}

	log.Println("Starting server on port :3000")
	server.ListenAndServe()
}
