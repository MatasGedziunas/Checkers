package models

import (
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
		logger.Fatalf("Invalid color given: %v ; expected w or b", color)
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
	// disableLogging()
	coppiedBoard := NewBoard(board.EncodeBoard())
	return checker.getCheckerPossibleMoves(coppiedBoard)
}

func (checker Tile) getCheckerPossibleMoves(board Board) []PossibleMove {
	if checker.isEmpty {
		return []PossibleMove{}
	}
	var directionToMoveRow int
	if checker.color == White {
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
			possibleCaptures := checker.getCheckerCaptures(board, checkerToCapture, NewEmptyTile(0, 0), []PossibleMove{}, 0)
			if len(possibleCaptures) > 0 {
				// logger.Printf("Found captures in checker: %v ; possible capture: %v ; length of possibleCaptures: %v", checker, checkerToCapture, len(possibleCaptures))
				captures = append(captures, possibleCaptures...)
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
		var filteredMoves []PossibleMove
		for _, move := range captures {
			if move.CapturesCount == maxCapturesCount {
				filteredMoves = append(filteredMoves, move)
			}
		}
		return filteredMoves
	} else {
		return checker.getNonCaptureMoves(board, directionToMoveRow)
	}
}

func (checker *Tile) getCheckerCaptures(board Board, checkerToCapture Tile, prevCapture Tile, possibleMoves []PossibleMove, captureCount int) []PossibleMove {
	logger.Printf("Checker: %v ; Checking if can capture: %v", checker, checkerToCapture)
	if checker.CanCapture(board, checkerToCapture, prevCapture) {
		captureCount += 1
		board.GetChecker(checkerToCapture.cords.Row, checkerToCapture.cords.Col).SetEmpty()
		board.GetChecker(checker.cords.Row, checker.cords.Col).SetEmpty()
		movesAfterCapture := checker.getMovesAfterCapture(&board, &checkerToCapture)
		logger.Printf("Can capture, moves after Capture: %v \n", movesAfterCapture)
		logger.Printf("Possible moves: %v \n ", possibleMoves)
		addPossibleMoves(&possibleMoves, movesAfterCapture, captureCount)
		directions := []int{-1, 1}
		for _, move := range movesAfterCapture {
			for _, rowDirection := range directions {
				for _, colDirection := range directions {
					tileAfterCapture := board.GetChecker(move.Row, move.Col)
					var nextCapture *Tile
					if checker.isQueen {
						tileAfterCapture.SetQueen()
						nextCapture = tileAfterCapture.getFirstCheckerInDirectionForQueen(&board, rowDirection, colDirection)
					} else {
						board.GetChecker(move.Row, move.Col).SetChecker()
						nextCapture = board.GetChecker(
							tileAfterCapture.cords.Row+rowDirection,
							tileAfterCapture.cords.Col+colDirection,
						)
					}
					possibleMoves =
						tileAfterCapture.
							getCheckerCaptures(board, *nextCapture, checkerToCapture, possibleMoves, captureCount)
					board.GetChecker(move.Row, move.Col).SetEmpty()
					possibleMoves = removeNotMaxCaptureCountPossibleMoves(possibleMoves)
				}
			}
		}
	}
	return possibleMoves
}

func addPossibleMoves(possibleMoves *[]PossibleMove, movesAfterCapture []Coordinates, capturesCount int) {
	for _, cords := range movesAfterCapture {
		*possibleMoves = append(*possibleMoves, NewPossibleMove(cords.Row, cords.Col, capturesCount))
	}
}

func removeNotMaxCaptureCountPossibleMoves(possibleMoves []PossibleMove) []PossibleMove {
	capturesCount := 0
	for _, move := range possibleMoves {
		capturesCount = max(move.CapturesCount, capturesCount)
	}
	newPossibleMoves := []PossibleMove{}
	for _, move := range possibleMoves {
		if move.CapturesCount == capturesCount {
			newPossibleMoves = append(newPossibleMoves, move)
		}
	}
	return newPossibleMoves
}

func (queen *Tile) getQueenPossibleMoves(board Board) []PossibleMove {
	possibleMoves := []PossibleMove{}
	return possibleMoves
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
	logger.Printf("CanCapture: directionRow: %v, directionCol %v, prevDirectionRow %v, prevDirectionCol %v", directionRow, directionCol, prevDirectionRow, prevDirectionCol)
	if !checkerFrom.isQueen {
		return checkCaptureConditions(board, *checkerFrom, checkerTo, directionRow, directionCol)
	}
	if !checkerFrom.isSameDiagonal(&checkerTo) || ((prevDirectionRow*-1) == directionRow && (prevDirectionCol*-1) == directionCol) { // different row direction and col direction {
		return false
	}
	checkerToCheck := checkerFrom.getFirstCheckerInDirectionForQueen(&board, directionRow, directionCol)
	logger.Printf("CheckerToCheck: %v", checkerToCheck)
	return checkCaptureConditions(board, *checkerFrom, *checkerToCheck, directionRow, directionCol)
}

func checkCaptureConditions(board Board, checkerFrom, checkerTo Tile, directionRow, directionCol int) bool {
	return (utils.IsInBounds(checkerTo.cords.Row+directionRow, checkerTo.cords.Col+directionCol, len(board.Pieces)) &&
		!checkerFrom.isEmpty && !checkerTo.isEmpty &&
		checkerFrom.color != checkerTo.color &&
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

func (checker *Tile) getNonCaptureMoves(board Board, rowDirection int) []PossibleMove {
	possibleMoves := []PossibleMove{}
	if !checker.isQueen {
		if checker.canMoveToTile(board, checker.cords.Row+rowDirection, checker.cords.Col-1) {
			possibleMoves = append(possibleMoves, NewPossibleMove(checker.cords.Row+rowDirection, checker.cords.Col-1, 0))
		}
		if checker.canMoveToTile(board, checker.cords.Row+rowDirection, checker.cords.Col+1) {
			possibleMoves = append(possibleMoves, NewPossibleMove(checker.cords.Row+rowDirection, checker.cords.Col+1, 0))
		}

	} else {
		directions := []int{-1, 1}
		for _, rowD := range directions {
			for _, colD := range directions {
				curRow := checker.cords.Row + rowD
				curCol := checker.cords.Col + colD
				for utils.IsInBounds(curRow, curCol, len(board.Pieces)) {
					if checker.canMoveToTile(board, curRow, curCol) {
						possibleMoves = append(possibleMoves, NewPossibleMove(curRow, curCol, 0))
					} else {
						break
					}
					curRow += rowD
					curCol += colD
				}
			}
		}
	}
	return possibleMoves
}

func (checker *Tile) canMoveToTile(board Board, row, col int) bool {
	return utils.IsInBounds(row, col, board.boardSize) && board.GetChecker(row, col).isEmpty
}

func (checker *Tile) isSameDiagonal(checkerTo *Tile) bool {
	return math.Abs(float64(checker.cords.Row-checkerTo.cords.Row)) == math.Abs(float64(checker.cords.Col-checkerTo.cords.Col))
}
