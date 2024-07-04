package gameFunctionality

import (
	"net/http"

	"github.com/MatasGedziunas/Checkers.git/models"
)

func GetStartingBoard() string {
	return StartingBoard
}

func GetCapturesBoard() string {
	return CapturesBoard
}

func GetTripleCapturesBoard() string {
	return TripleNotDoubleCaptureBoard
}

func DecodeBoard(boardString string) models.Board {
	return models.NewBoard(boardString)
}

func EncodeBoard(board models.Board) string {
	return board.EncodeBoard()
}

func GetPossibleMoves(w http.ResponseWriter, r *http.Request) {
	boardString := r.URL.Query().Get("boardString")
	row := r.URL.Query().Get("row")
	col := r.URL.Query().Get("col")
	if (row == "" && col != "") || (row != "" && col == "") {
		http.Error(w, "Neither or both row and col should be given", http.StatusBadRequest)
	}
}
