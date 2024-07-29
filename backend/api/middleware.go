package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/utils"
)

func withParsedBody[T utils.ValidatedModel](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.ReadBody[T](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c := context.WithValue(context.Background(), models.GameInfoKey, body)
		next.ServeHTTP(w, r.WithContext(c))
	}
}

func withDB(db *sql.DB, handler func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	}
}

func dbErrorMiddleware(dbError error) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if dbError != nil {
				http.Error(w, fmt.Sprintf("Database connection error: %v", dbError), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
