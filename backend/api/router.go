package main

import (
	"fmt"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
)

func main() {
	board := models.NewBoard(gameFunctionality.GetCapturesBoard())
	fmt.Println(board.EncodeBoard())
	// for _, row := range board.Pieces {
	// 	for _, tile := range row {
	// 		possibleMoves := tile.GetPossibleMoves(board)
	// 		if len(possibleMoves) > 0 {
	// 			fmt.Printf("Found captures Possible move: %v ; for checker: %v \n", possibleMoves, tile)
	// 			fmt.Println(board.EncodeBoard())
	// 		}
	// 	}
	// }
	checker70 := board.GetChecker(7, 2)
	fmt.Printf("%v ", checker70.GetPossibleMoves(board))
}
