package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

const defaultBoardSize = 10

type Board struct {
	Pieces    [][]Tile
	boardSize int
}

func NewBoard(boardString string) (Board, error) {
	board := strings.Split(boardString, " ")
	if len(board) != defaultBoardSize {
		return Board{}, fmt.Errorf("invalid board given, expected size to be %v, but got: %v ; board: %v ; boardString: %v", defaultBoardSize, len(board), board, boardString)
	}
	var pieces [][]Tile
	var row []Tile
	curRow := 0
	for curRow < len(board) {
		row = []Tile{}
		curStr := 0
		for curStr < len(board[curRow]) {
			var piece Tile
			curCol := len(row)
			curColor := string(board[curRow][curStr])
			if board[curRow][curStr] == '.' {
				piece = NewEmptyTile(curRow, curCol)
			} else if len(board[curRow]) > curStr+1 && board[curRow][curStr+1] == 'q' {
				piece = NewQueen(curRow, curCol, curColor)
				curStr += 1
			} else {
				piece = NewChecker(curRow, curCol, curColor)
			}
			row = append(row, piece)
			curStr += 1
		}
		if len(row) != defaultBoardSize {
			return Board{}, fmt.Errorf("invalid board given, expected row size to be %v, but got: %v ; row: %v ; board: %v", defaultBoardSize, curStr, board[curRow], board)
		}
		pieces = append(pieces, row)
		curRow += 1
	}
	return Board{Pieces: pieces,
		boardSize: len(row)}, nil
}

func (board *Board) EncodeBoard() (string, error) {
	var sb strings.Builder
	if len(board.Pieces) != defaultBoardSize {
		return "", fmt.Errorf("Encode board, Invalid row size: %v ; expected %v", len(board.Pieces), defaultBoardSize)
	}
	for rowIndex, row := range board.Pieces {
		for _, piece := range row {
			if piece.isEmpty {
				sb.WriteByte('.')
			} else {
				sb.WriteString(piece.Color.GetColorString())
			}
			if piece.isQueen {
				sb.WriteByte('q')
			}
		}
		if rowIndex != defaultBoardSize-1 {
			sb.WriteByte(' ')
		}

	}
	return sb.String(), nil
}

func (board *Board) PrintBoard() {
	for _, row := range board.Pieces {
		rowToPrint := ""
		for _, piece := range row {
			rowToPrint += piece.ToStr()
		}
		log.Print(rowToPrint)
	}
}

func (board *Board) GetChecker(row int, col int) *Tile {
	if utils.IsInBounds(row, col, board.boardSize) {
		return &board.Pieces[row][col]
	} else {
		// log.Printf("Trying to get out of bounds checker: row %v ; col %v", row, col)
		tempTile := NewEmptyTile(0, 0)
		return &tempTile
	}
}
