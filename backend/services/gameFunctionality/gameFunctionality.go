package gameFunctionality

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func DecodeBoard(boardString string) (models.Board, error) {
	return models.NewBoard(boardString)
}

func EncodeBoard(board models.Board) (string, error) {

	return board.EncodeBoard()
}

func GetPossibleMoves(w http.ResponseWriter, r *http.Request) {
	boardString := r.URL.Query().Get("boardString")
	rowStr := r.URL.Query().Get("row")
	colStr := r.URL.Query().Get("col")
	if boardString == "" {
		http.Error(w, "boardString query param not given", http.StatusBadRequest)
		return
	}
	if (rowStr == "" && colStr != "") || (rowStr != "" && colStr == "") {
		http.Error(w, "Neither or both row and col should be given", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	board, err := DecodeBoard(boardString)
	board.PrintBoard()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if rowStr != "" {
		row, err := strconv.Atoi(rowStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Problem parsing row to int, row: %v", rowStr), http.StatusBadRequest)
			return
		}
		col, err := strconv.Atoi(colStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Problem parsing col to int, col: %v", colStr), http.StatusBadRequest)
			return
		}
		p, err := board.GetChecker(row, col).GetPossibleMoves(board)
		if err != nil {
			http.Error(w, fmt.Sprintf("Problem getting possible moves for checker %v ; Error: %v", board.GetChecker(row, col), err.Error()), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(p)
	} else {
		possibleMoves := make([][]([]models.PossibleMove), len(board.Pieces))
		for rowIndex, row := range board.Pieces {
			for _, piece := range row {
				p, err := piece.GetPossibleMoves(board)
				if err != nil {
					http.Error(w, fmt.Sprintf("Problem getting possible moves for checker %v ; Error: %v", piece, err.Error()), http.StatusInternalServerError)
					return
				}
				possibleMoves[rowIndex] = append(possibleMoves[rowIndex], p)
			}
		}
		json.NewEncoder(w).Encode(possibleMoves)
	}
}
