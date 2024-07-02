package gameFunctionality

import "github.com/MatasGedziunas/Checkers.git/models"

func GetStartingBoard() string {
	return StartingBoard
}

func GetCapturesBoard() string {
	return CapturesBoard
}

func DecodeBoard(boardString string) models.Board {
	return models.NewBoard(boardString)
}

func EncodeBoard(board models.Board) string {
	return board.EncodeBoard()
}
