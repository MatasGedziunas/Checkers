package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/services/databaseFunctionality"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func withParsedBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyInBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Problem reading body of request: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		var body databaseFunctionality.Game
		err = json.Unmarshal(bodyInBytes, &body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Problem parsing body in json: %s", err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Printf("Parsed body: %+v", body)
		c := context.WithValue(context.Background(), models.GameInfoKey, body)
		next.ServeHTTP(w, r.WithContext(c))
	}
}

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
	r.Use(cors.Handler(corsOptions))
	r.Get("/possibleMoves", gameFunctionality.GetPossibleMoves)
	r.Post("/saveBoard", withParsedBody(func(w http.ResponseWriter, r *http.Request) {
		db, err := connectDB()
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to establish connection with database: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		databaseFunctionality.StoreGame(w, r, db)
	}))

	log.Println("Server running")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
