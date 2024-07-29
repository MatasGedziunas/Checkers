package databaseFunctionality

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

func StoreGame(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("In store game function")
	game, err := utils.ReadBody[Game](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Game data: %+v\n", game)
	_, err = db.Query("INSERT INTO games (white_player, black_player) VALUES ($1, $2)", game.WhitePlayer, game.BlackPlayer)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem inserting record into database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Game data stored successfully"))
}

func StoreUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("In store user function")
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

func StoreBoard(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("In store boards function")
	board, err := utils.ReadBody[Board](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Board data: %+v\n", board)
	_, err = db.Query("INSERT INTO boards (board_string, game_id) VALUES ($1, $2)", board.BoardString, board.GameId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem inserting record into database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Board data stored successfully"))
}
