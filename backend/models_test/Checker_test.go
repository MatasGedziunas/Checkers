package models_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/MatasGedziunas/Checkers.git/models"
	"github.com/MatasGedziunas/Checkers.git/services/gameFunctionality"
)

func TestChecker_GetPossibleMoves(t *testing.T) {
	capturesBoard, _ := gameFunctionality.DecodeBoard(gameFunctionality.GetCapturesBoard())
	queenCaptureBoard, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenCaptureSimple)
	queenTripleCaptureBoard, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenTripleTrickyCapture)
	queenDifficultCapture, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenDifficultCapture)
	queenStopWithSameColorChecker, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenStopCaptureWithSameColorChecker)
	t.Run("OneCapture", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(6, 3)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(capturesBoard)
		want := []models.PossibleMove{models.NewPossibleMove(4, 5, 1)}
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
	t.Run("TwoCaptures", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(capturesBoard)
		want := []models.PossibleMove{models.NewPossibleMove(5, 0, 2)}
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenDoubleCapture", func(t *testing.T) {
		checkerToTest := queenCaptureBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenCaptureBoard)
		want := []models.PossibleMove{models.NewPossibleMove(7, 8, 2), models.NewPossibleMove(8, 9, 2)}
		sortPossibleMoves(possibleMoves)
		sortPossibleMoves(want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenTrickyCapture", func(t *testing.T) {
		checkerToTest := queenTripleCaptureBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenTripleCaptureBoard)
		want := []models.PossibleMove{models.NewPossibleMove(7, 8, 2), models.NewPossibleMove(8, 9, 2), models.NewPossibleMove(2, 3, 2), models.NewPossibleMove(1, 2, 2), models.NewPossibleMove(0, 1, 2)}
		sortPossibleMoves(possibleMoves)
		sortPossibleMoves(want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenDifficultCapture", func(t *testing.T) {
		checkerToTest := queenDifficultCapture.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenDifficultCapture)
		want := []models.PossibleMove{models.NewPossibleMove(7, 8, 3), models.NewPossibleMove(8, 9, 3), models.NewPossibleMove(0, 9, 3), models.NewPossibleMove(1, 2, 3), models.NewPossibleMove(0, 1, 3)}
		sortPossibleMoves(possibleMoves)
		sortPossibleMoves(want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenStopCaptureWithSameColorChecker", func(t *testing.T) {
		checkerToTest := queenStopWithSameColorChecker.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenStopWithSameColorChecker)
		want := []models.PossibleMove{models.NewPossibleMove(6, 3, 1)}
		sortPossibleMoves(possibleMoves)
		sortPossibleMoves(want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

}

func sortPossibleMoves(moves []models.PossibleMove) {
	sort.Slice(moves, func(i, j int) bool {
		if moves[i].CapturesCount != moves[j].CapturesCount {
			return moves[i].CapturesCount < moves[j].CapturesCount
		}
		if moves[i].Move.Row != moves[j].Move.Row {
			return moves[i].Move.Row < moves[j].Move.Row
		}
		return moves[i].Move.Col < moves[j].Move.Col
	})
}
