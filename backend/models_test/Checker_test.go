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
	queenCaptureBoard, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenDoubleCapture)
	queenTripleCaptureBoard, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenTrickyCapture)
	queenDifficultCapture, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenDifficultCapture)
	queenStopWithSameColorChecker, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenStopCaptureWithSameColorChecker)
	queenWeirdStuff, _ := gameFunctionality.DecodeBoard(gameFunctionality.QueenWeirdStuff)
	t.Run("OneCapture", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(6, 3)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(capturesBoard)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(4, 5)}
		want := models.NewPossibleMove(1, wantCoordinates)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
	t.Run("TwoCaptures", func(t *testing.T) {
		checkerToTest := capturesBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(capturesBoard)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(7, 2)}
		want := models.NewPossibleMove(2, wantCoordinates)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenDoubleCapture", func(t *testing.T) {
		checkerToTest := queenCaptureBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenCaptureBoard)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(4, 5)}
		want := models.NewPossibleMove(2, wantCoordinates)
		sortPossibleMoves(&possibleMoves)
		sortPossibleMoves(&want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenTrickyCapture", func(t *testing.T) {
		checkerToTest := queenTripleCaptureBoard.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenTripleCaptureBoard)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(4, 5)}
		want := models.NewPossibleMove(2, wantCoordinates)
		sortPossibleMoves(&possibleMoves)
		sortPossibleMoves(&want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenDifficultCapture", func(t *testing.T) {
		checkerToTest := queenDifficultCapture.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenDifficultCapture)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(6, 3)}
		want := models.NewPossibleMove(3, wantCoordinates)
		sortPossibleMoves(&possibleMoves)
		sortPossibleMoves(&want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})

	t.Run("QueenStopCaptureWithSameColorChecker", func(t *testing.T) {
		checkerToTest := queenStopWithSameColorChecker.GetChecker(9, 0)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenStopWithSameColorChecker)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(6, 3)}
		want := models.NewPossibleMove(1, wantCoordinates)
		sortPossibleMoves(&possibleMoves)
		sortPossibleMoves(&want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
	t.Run("QueenWeirdStuff", func(t *testing.T) {
		checkerToTest := queenWeirdStuff.GetChecker(8, 7)
		possibleMoves, _ := checkerToTest.GetPossibleMoves(queenWeirdStuff)
		wantCoordinates := []models.Coordinates{models.NewCoordinates(5, 4)}
		want := models.NewPossibleMove(2, wantCoordinates)
		sortPossibleMoves(&possibleMoves)
		sortPossibleMoves(&want)
		if !reflect.DeepEqual(possibleMoves, want) {
			t.Errorf("%v checker tested ; got possibleMoves: %v ; expected: %v", checkerToTest, possibleMoves, want)
		}
	})
}

func sortPossibleMoves(move *models.PossibleMove) {
	sort.Slice(move.Moves, func(i, j int) bool {
		if move.Moves[i].Row != move.Moves[j].Row {
			return move.Moves[i].Row < move.Moves[j].Row
		}
		return move.Moves[i].Col < move.Moves[j].Col
	})
}
