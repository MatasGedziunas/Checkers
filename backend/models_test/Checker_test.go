package models_test

import (
	"reflect"
	"testing"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
)

func TestChecker_GetPossibleMoves(t *testing.T) {
	capturesBoard := gameFunctionality.DecodeBoard(gameFunctionality.GetCapturesBoard())
	t.Run("OneCapture", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(6, 3)
		possibleMoves := checkerToTest.GetPossibleMoves(capturesBoard)
		want := []models.PossibleMove{models.NewPossibleMove(5, 4, 1)}
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
	t.Run("TwoCaptures", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(9, 0)
		possibleMoves := checkerToTest.GetPossibleMoves(capturesBoard)
		want := []models.PossibleMove{models.NewPossibleMove(8, 1, 2)}
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
}
