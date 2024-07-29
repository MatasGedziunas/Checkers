package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	authentication "github.com/MatasGedziunas/Checkers.git/services/authenticationFunctionality"
	"github.com/MatasGedziunas/Checkers.git/services/databaseFunctionality"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func connectDB() (*sql.DB, error) {
	password := os.Getenv("postgresql_password")
	connStr := fmt.Sprintf("user=postgres dbname=checkers sslmode=disable password=%s", password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
func main() {
	r := chi.NewRouter()
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
	}
	db, dbError := connectDB()
	r.Use(cors.Handler(corsOptions))
	r.Get("/possibleMoves", gameFunctionality.GetPossibleMoves)
	r.Route("/db", func(r chi.Router) {
		r.Use(dbErrorMiddleware(dbError))
		r.Post("/StoreGame", withDB(db, databaseFunctionality.StoreGame))
		r.Post("/StoreUser", withDB(db, databaseFunctionality.StoreUser))
		r.Post("/StoreBoard", withDB(db, databaseFunctionality.StoreBoard))

		r.Get("/GetGamesOfUser", withDB(db, databaseFunctionality.GetGamesOfUser))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", withDB(db, authentication.Login))
	})

	log.Println("Server running")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
