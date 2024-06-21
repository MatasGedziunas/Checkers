package main

import (
	"fmt"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
)

func main() {
	board := models.NewBoard(gameFunctionality.GetStartingBoard())
	fmt.Println(board.EncodeBoard())
}
