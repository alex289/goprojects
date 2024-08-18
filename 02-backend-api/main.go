package main

import (
	"calcapi/db"
	"calcapi/handler"
	"calcapi/middleware"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	err := db.InitDb()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /add", handler.AddHandler)
	router.HandleFunc("POST /subtract", handler.SubtractHandler)
	router.HandleFunc("POST /multiply", handler.MultiplyHandler)
	router.HandleFunc("POST /divide", handler.DivideHandler)
	router.HandleFunc("POST /sum", handler.SumHandler)

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Headers,
		middleware.RateLimit,
		middleware.Identify,
	)

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(cors.Default().Handler(router)),
	}

	log.Println("Starting server on port :3000")
	server.ListenAndServe()
}
