package models

import (
	"errors"
	"log"
	"math"
	"os"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	disableLogging()
}

func enableLogging() {
	logger.SetOutput(os.Stdout)
}

func disableLogging() {
	logger.SetOutput(os.Stderr)
}

type Tile struct {
	cords   Coordinates
	isEmpty bool
	isQueen bool
	Color   Color
}

func NewChecker(row int, col int, color string) Tile {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	var tempColor Color = getColorFromString(color)
	return Tile{cords: cords,
		Color: tempColor,
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
		Color:   tempColor,
	}
}

func getColorFromString(color string) Color {
	if color == "w" {
		return White
	} else if color == "b" {
		return Black
	} else {
		logger.Fatalf("Invalid color given: %v ; expected w or b", color)
		return ""
	}
}

func (checker *Tile) ToStr() string {
	str := ""
	if checker.isEmpty {
		str = "."
	} else {
		str = string(checker.Color)
	}
	if checker.isQueen {
		str += "q"
	}
	return str
}

func (checker *Tile) GetCoordinates() Coordinates {
	return checker.cords
}

func (checker *Tile) SetEmpty() {
	checker.isEmpty = true
	checker.isQueen = false
	checker.Color = ""
}

func (checker *Tile) SetChecker() {
	checker.isEmpty = false
}

func (checker *Tile) SetOptions(isEmpty bool, isQueen bool, color Color) {
	checker.isEmpty = isEmpty
	checker.isQueen = isQueen
	checker.Color = color
}

func (checker *Tile) SetQueen() {
	checker.isEmpty = false
	checker.isQueen = true
}

func (checker *Tile) GetPossibleMoves(board Board) (PossibleMove, error) {
	disableLogging()
	encodedBoard, err := board.EncodeBoard()
	if err != nil {
		return PossibleMove{}, errors.New(err.Error())
	}
	coppiedBoard, err := NewBoard(encodedBoard)
	if err != nil {
		log.Fatal(err.Error())
	}
	return checker.getCheckerPossibleMoves(coppiedBoard), nil
}

func (checker Tile) getCheckerPossibleMoves(board Board) PossibleMove {
	if checker.isEmpty {
		return PossibleMove{}
	}
	var directionToMoveRow int
	if checker.Color == White {
		directionToMoveRow = -1
	} else {
		directionToMoveRow = 1
	}
	captures := []PossibleMove{}
	directions := [2]int{-1, 1}
	for _, rowDirection := range directions {
		for _, colDirection := range directions {
			var checkerToCapture Tile = *board.GetChecker(checker.cords.Row+rowDirection, checker.cords.Col+colDirection)
			if checker.isQueen {
				checkerToCapture = *checker.getFirstCheckerInDirectionForQueen(&board, rowDirection, colDirection)
			}
			capturesCount, possibleCaptures := checker.getCheckerCaptures(board, checkerToCapture, NewEmptyTile(0, 0), []Coordinates{}, 0)
			if len(possibleCaptures) > 0 {
				// logger.Printf("Found captures in checker: %v ; possible capture: %v ; length of possibleCaptures: %v", checker, checkerToCapture, len(possibleCaptures))
				captures = append(captures, NewPossibleMove(capturesCount, possibleCaptures))
				// need to append checker that was captured first because this is sent to frontend
				// this can be improved by sending all steps of a capture...
			}
		}
	}
	logger.Printf("Captures for checker %v : %v", checker, captures)
	if len(captures) > 0 {
		// get max capture count moves
		var maxCapturesCount int
		for _, move := range captures {
			maxCapturesCount = max(maxCapturesCount, move.CapturesCount)
		}
		var filteredMoves []Coordinates
		for _, move := range captures {
			if move.CapturesCount == maxCapturesCount {
				filteredMoves = append(filteredMoves, move.Moves...)
			}
		}
		return NewPossibleMove(maxCapturesCount, filteredMoves)
	} else {
		return checker.getNonCaptureMoves(board, directionToMoveRow)
	}
}

func (checker *Tile) getCheckerCaptures(board Board, checkerToCapture Tile, prevCapture Tile, possibleMoves []Coordinates, captureCount int) (int, []Coordinates) {
	logger.Printf("Checker: %v ; Checking if can capture: %v", checker, checkerToCapture)
	if checker.CanCapture(board, checkerToCapture, prevCapture) {
		captureCount += 1
		// logger.Printf("Checker: %v ; can capture: %v ; capturesCount: %v", checker, checkerToCapture, captureCount)
		board.GetChecker(checkerToCapture.cords.Row, checkerToCapture.cords.Col).SetEmpty()
		board.GetChecker(checker.cords.Row, checker.cords.Col).SetEmpty()
		var movesAfterCapture []Coordinates = checker.getMovesAfterCapture(&board, &checkerToCapture)
		// logger.Printf("Can capture, moves after Capture: %v \n", movesAfterCapture)
		directions := []int{-1, 1}
		startingCaptureCount := captureCount
		for _, move := range movesAfterCapture {
			var tempCaptureCount int
			for _, rowDirection := range directions {
				for _, colDirection := range directions {
					tileAfterCapture := board.GetChecker(move.Row, move.Col)
					var nextCapture *Tile
					if checker.isQueen {
						tileAfterCapture.SetOptions(false, true, checker.Color)
						nextCapture = tileAfterCapture.getFirstCheckerInDirectionForQueen(&board, rowDirection, colDirection)
					} else {
						tileAfterCapture.SetOptions(false, false, checker.Color)
						nextCapture = board.GetChecker(
							tileAfterCapture.cords.Row+rowDirection,
							tileAfterCapture.cords.Col+colDirection,
						)
					}
					// logger.Printf("RowDirection: %v, colDirection: %v ; nextCapture: %v, tileAfterCapture: %v, checker: %v", rowDirection, colDirection, nextCapture, tileAfterCapture, checker)
					var temp int
					temp, possibleMoves =
						tileAfterCapture.
							getCheckerCaptures(board, *nextCapture, checkerToCapture, possibleMoves, startingCaptureCount)
					tileAfterCapture.SetEmpty()
					tempCaptureCount = max(tempCaptureCount, temp)
				}
			}
			if startingCaptureCount == 1 && tempCaptureCount == captureCount {
				possibleMoves = append(possibleMoves, NewCoordinates(move.Row, move.Col))
			}
			if startingCaptureCount == 1 && tempCaptureCount > captureCount {
				possibleMoves = []Coordinates{NewCoordinates(move.Row, move.Col)}
			}
			captureCount = max(captureCount, tempCaptureCount)
		}
	}
	return captureCount, possibleMoves
}

func (checker *Tile) getMovesAfterCapture(board *Board, capturedChecker *Tile) []Coordinates {
	directionRow := utils.GetDirection(checker.cords.Row, capturedChecker.cords.Row)
	directionCol := utils.GetDirection(checker.cords.Col, capturedChecker.cords.Col)
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
			curRow += directionRow
			curCol += directionCol
		}
	}
	return moves
}

func (checkerFrom *Tile) CanCapture(board Board, checkerTo Tile, prevChecker Tile) bool {
	directionRow := utils.GetDirection(checkerFrom.cords.Row, checkerTo.cords.Row)
	directionCol := utils.GetDirection(checkerFrom.cords.Col, checkerTo.cords.Col)
	prevDirectionRow := utils.GetDirection(prevChecker.cords.Row, checkerFrom.cords.Row)
	prevDirectionCol := utils.GetDirection(prevChecker.cords.Col, checkerFrom.cords.Col)
	// logger.Printf("CanCapture: directionRow: %v, directionCol %v, prevDirectionRow %v, prevDirectionCol %v", directionRow, directionCol, prevDirectionRow, prevDirectionCol)
	if !checkerFrom.isQueen {
		return checkCaptureConditions(board, *checkerFrom, checkerTo, directionRow, directionCol)
	}
	// logger.Print(
	// 	"isSameDiagonal: ", checkerFrom.isSameDiagonal(&checkerTo),
	// 	", prevChecker.isDummyChecker: ", prevChecker.isDummyChecker(),
	// 	", prevDirectionRow: ", prevDirectionRow, ", directionRow: ", directionRow,
	// 	", prevDirectionCol: ", prevDirectionCol, ", directionCol: ", directionCol,
	// )
	isSameDiagonal := checkerFrom.isSameDiagonal(&checkerTo)
	sameDirectionAsPrevCapturedChecker := (prevDirectionRow*-1) == directionRow && (prevDirectionCol*-1) == directionCol
	if prevChecker.isDummyChecker() {
		sameDirectionAsPrevCapturedChecker = false
	}

	// logger.Print("isSameDiagonal: ", isSameDiagonal, ", sameDirection: ", sameDirectionAsPrevCapturedChecker)

	if !isSameDiagonal || sameDirectionAsPrevCapturedChecker {
		return false
	}
	checkerToCheck := checkerFrom.getFirstCheckerInDirectionForQueen(&board, directionRow, directionCol)
	// logger.Printf("CheckerToCheck: %v", checkerToCheck)
	return checkCaptureConditions(board, *checkerFrom, *checkerToCheck, directionRow, directionCol)
}

func checkCaptureConditions(board Board, checkerFrom, checkerTo Tile, directionRow, directionCol int) bool {
	return (utils.IsInBounds(checkerTo.cords.Row+directionRow, checkerTo.cords.Col+directionCol, len(board.Pieces)) &&
		!checkerFrom.isEmpty && !checkerTo.isEmpty &&
		checkerFrom.Color != checkerTo.Color &&
		board.GetChecker(checkerTo.cords.Row+directionRow, checkerTo.cords.Col+directionCol).isEmpty &&
		checkerFrom.isSameDiagonal(&checkerTo))
}

func (queen *Tile) getFirstCheckerInDirectionForQueen(board *Board, rowDirection, colDirection int) *Tile {
	if !queen.isQueen {
		logger.Fatal("getFirstCheckerInDirectionForQueen called when tile is not queen")
	}
	curRow := queen.cords.Row + rowDirection
	curCol := queen.cords.Col + colDirection
	for utils.IsInBounds(curRow, curCol, board.boardSize) {

		if !board.GetChecker(curRow, curCol).isEmpty {
			return board.GetChecker(curRow, curCol)
		}
		curRow += rowDirection
		curCol += colDirection
	}
	return &Tile{isEmpty: true}
}

func (checker *Tile) isDummyChecker() bool {
	return checker.cords.Row == 0 && checker.cords.Col == 0
}

func (checker *Tile) getNonCaptureMoves(board Board, rowDirection int) PossibleMove {
	possibleMoves := []Coordinates{}
	if !checker.isQueen {
		if checker.canMoveToTile(board, checker.cords.Row+rowDirection, checker.cords.Col-1) {
			possibleMoves = append(possibleMoves, NewCoordinates(checker.cords.Row+rowDirection, checker.cords.Col-1))
		}
		if checker.canMoveToTile(board, checker.cords.Row+rowDirection, checker.cords.Col+1) {
			possibleMoves = append(possibleMoves, NewCoordinates(checker.cords.Row+rowDirection, checker.cords.Col+1))
		}
		return NewPossibleMove(0, possibleMoves)
	} else {
		directions := []int{-1, 1}
		for _, rowD := range directions {
			for _, colD := range directions {
				curRow := checker.cords.Row + rowD
				curCol := checker.cords.Col + colD
				for utils.IsInBounds(curRow, curCol, len(board.Pieces)) {
					if checker.canMoveToTile(board, curRow, curCol) {
						possibleMoves = append(possibleMoves, NewCoordinates(curRow, curCol))
					} else {
						break
					}
					curRow += rowD
					curCol += colD
				}
			}
		}
	}
	return NewPossibleMove(0, possibleMoves)
}

func (checker *Tile) canMoveToTile(board Board, row, col int) bool {
	return utils.IsInBounds(row, col, board.boardSize) && board.GetChecker(row, col).isEmpty
}

func (checker *Tile) isSameDiagonal(checkerTo *Tile) bool {
	return math.Abs(float64(checker.cords.Row-checkerTo.cords.Row)) == math.Abs(float64(checker.cords.Col-checkerTo.cords.Col))
}
