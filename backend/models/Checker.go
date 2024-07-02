package models

import (
	"log"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

type Tile struct {
	cords   Coordinates
	isEmpty bool
	isQueen bool
	color   Color
}

func NewChecker(row int, col int, color string) Tile {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	var tempColor Color = getColorFromString(color)
	return Tile{cords: cords,
		color: tempColor,
	}
}

func NewEmptyTile(row int, col int) Tile {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	return Tile{cords: cords,
		isEmpty: true}
}

func NewQueen(row int, col int, color string) Tile {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	var tempColor Color = getColorFromString(color)
	return Tile{cords: cords,
		isQueen: true,
		color:   tempColor,
	}
}

func getColorFromString(color string) Color {
	if color == "w" {
		return White
	} else if color == "b" {
		return Black
	} else {
		log.Fatalf("Invalid color given: %v ; expected w or b", color)
		return ""
	}
}

func (checker *Tile) GetCoordinates() Coordinates {
	return checker.cords
}

func (checker *Tile) SetEmpty() {
	checker.isEmpty = true
}

func (checker *Tile) SetChecker() {
	checker.isEmpty = false
}

func (checker *Tile) SetQueen() {
	checker.isEmpty = false
	checker.isQueen = true
}

func (checker *Tile) GetPossibleMoves(board Board) []PossibleMove {
	coppiedBoard := NewBoard(board.EncodeBoard())
	return checker.getCheckerPossibleMoves(coppiedBoard)
}

func (checker Tile) getCheckerPossibleMoves(board Board) []PossibleMove {
	if checker.isEmpty {
		return []PossibleMove{}
	}
	// var directionToMove int
	// if checker.color == White {
	// 	directionToMove = -1
	// } else {
	// 	directionToMove = 1
	// }
	captures := []PossibleMove{}
	directions := [2]int{-1, 1}
	for _, rowDirection := range directions {
		for _, colDirection := range directions {
			var checkerToCapture Tile = *board.GetChecker(checker.cords.Row+rowDirection, checker.cords.Col+colDirection)
			if checker.isQueen {
				checkerToCapture = *checker.getFirstCheckerInDirectionForQueen(&board, rowDirection, colDirection)
			}
			capturesCount := checker.getCheckerCaptures(board, checkerToCapture, 0)
			if capturesCount > 0 {
				log.Printf("Found captures in checker: %v ; possible capture: %v", checker, checkerToCapture)
				// if queen need to check for all places where it can move
				// etc. getPossibleTilesAfterCapture
				captures = append(captures, NewPossibleMove(checkerToCapture.cords.Row, checkerToCapture.cords.Col, capturesCount))
			}
		}
	}

	if len(captures) > 0 {
		// get max capture count moves
		var maxCapturesCount int
		for _, move := range captures {
			maxCapturesCount = max(maxCapturesCount, move.CapturesCount)
		}
		var filteredMoves []PossibleMove
		for _, move := range captures {
			if move.CapturesCount == maxCapturesCount {
				filteredMoves = append(filteredMoves, move)
			}
		}
		return filteredMoves
	} else {

	}
	return []PossibleMove{}
}

func (checker *Tile) getCheckerCaptures(board Board, checkerToCapture Tile, capturesCount int) int {
	var maxCapturesCount int = capturesCount
	if checker.CanCapture(board, checkerToCapture) {
		capturesCount += 1
		maxCapturesCount = capturesCount
		board.GetChecker(checkerToCapture.cords.Row, checkerToCapture.cords.Col).SetEmpty()
		board.GetChecker(checker.cords.Row, checker.cords.Col).SetEmpty()
		movesAfterCapture := checker.getMovesAfterCapture(&board, &checkerToCapture)
		directions := []int{-1, 1}
		for _, move := range movesAfterCapture {
			for _, rowDirection := range directions {
				for _, colDirection := range directions {
					board.GetChecker(move.Row, move.Col).SetChecker()
					tileAfterCapture := board.GetChecker(move.Row, move.Col)
					tempCapturesCount :=
						tileAfterCapture.
							getCheckerCaptures(board, *board.GetChecker(
								tileAfterCapture.cords.Row+rowDirection,
								tileAfterCapture.cords.Col+colDirection,
							), capturesCount)
					maxCapturesCount = max(maxCapturesCount, tempCapturesCount)
					board.GetChecker(move.Row, move.Col).SetEmpty()
				}
			}
		}
	}
	return maxCapturesCount
}

func (queen *Tile) getQueenPossibleMoves(board Board) []PossibleMove {
	possibleMoves := []PossibleMove{}
	return possibleMoves
}

func (checker *Tile) getMovesAfterCapture(board *Board, capturedChecker *Tile) []Coordinates {
	directionRow := capturedChecker.cords.Row - checker.cords.Row
	directionCol := capturedChecker.cords.Col - checker.cords.Col
	moves := []Coordinates{}
	if !checker.isQueen { // Simple checker
		moves = append(moves, NewCoordinates(capturedChecker.cords.Row+directionRow, capturedChecker.cords.Col+directionCol))
	} else { // Queen
		curRow := capturedChecker.cords.Row + directionRow
		curCol := capturedChecker.cords.Col + directionCol
		for utils.IsInBounds(curRow, curCol, len(board.Pieces)) {
			if board.GetChecker(curRow, curCol).isEmpty {
				moves = append(moves, NewCoordinates(curRow, curCol))
			} else {
				break
			}
		}
	}
	return moves
}

func (checkerFrom *Tile) CanCapture(board Board, checkerTo Tile) bool {
	directionRow := checkerTo.cords.Row - checkerFrom.cords.Row
	directionCol := checkerTo.cords.Col - checkerFrom.cords.Col
	if !checkerFrom.isQueen {
		return checkCaptureConditions(board, *checkerFrom, checkerTo, directionRow, directionCol)
	}
	checkerToCheck := checkerFrom.getFirstCheckerInDirectionForQueen(&board, directionRow, directionCol)
	return checkCaptureConditions(board, *checkerFrom, *checkerToCheck, directionRow, directionCol)
}

func checkCaptureConditions(board Board, checkerFrom, checkerTo Tile, directionRow, directionCol int) bool {
	return (utils.IsInBounds(checkerTo.cords.Row+directionRow, checkerTo.cords.Col+directionCol, len(board.Pieces)) &&
		!checkerFrom.isEmpty && !checkerTo.isEmpty &&
		checkerFrom.color != checkerTo.color &&
		board.GetChecker(checkerTo.cords.Row+directionRow, checkerTo.cords.Col+directionCol).isEmpty)
}

func (queen *Tile) getFirstCheckerInDirectionForQueen(board *Board, rowDirection, colDirection int) *Tile {
	if !queen.isQueen {
		log.Fatal("getFirstCheckerInDirectionForQueen called when tile is not queen")
	}
	curRow := queen.cords.Row + rowDirection
	curCol := queen.cords.Col + colDirection
	for utils.IsInBounds(curRow, curCol, board.boardSize) {
		if !board.GetChecker(curRow, curCol).isEmpty {
			return board.GetChecker(curRow, curCol)
		}
	}
	return &Tile{isEmpty: true}
}
