package databaseFunctionality

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("In get user function")
	user, err := utils.ReadBody[User](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("User data: %+v\n", user)
	_, err = db.Query("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, utils.HashString(user.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem inserting record into database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User data stored successfully"))
}

func GetGamesOfUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("In get games for user function")
	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		http.Error(w, "No user_id parameter found in request url query", http.StatusBadRequest)
		return
	}
	fmt.Printf("user_id: %s\n", user_id)
	rows, err := db.Query("SELECT * FROM games WHERE white_player = $1", user_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem getting records from database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	var games []Game
	for rows.Next() {
		var game Game
		rows.Scan(utils.GetColumnsOfStruct(&game)...)
		fmt.Printf("row: %+v", game)
		games = append(games, game)
	}
	rows, err = db.Query("SELECT * FROM games WHERE black_player = $1", user_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem getting records from database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var game Game
		rows.Scan(utils.GetColumnsOfStruct(&game)...)
		fmt.Printf("row: %+v", game)
		games = append(games, game)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(games)
}
