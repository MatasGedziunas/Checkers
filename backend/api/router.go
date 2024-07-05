package main

import (
	"log"
	"net/http"

	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
	}
	r.Use(cors.Handler(corsOptions))
	r.Get("/possibleMoves", gameFunctionality.GetPossibleMoves)
	log.Println("Server running")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
