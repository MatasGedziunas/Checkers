package models

import (
	"strings"

	"github.com/MatasGedziunas/Checkers.git/utils"
)

type Board struct {
	Pieces    [][]Tile
	boardSize int
}

func NewBoard(boardString string) Board {
	board := strings.Split(boardString, " ")
	var pieces [][]Tile
	var row []Tile
	curRow := 0
	for curRow < len(board)-1 {
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
		pieces = append(pieces, row)
		curRow += 1
	}
	return Board{Pieces: pieces,
		boardSize: len(row)}
}

func (board *Board) EncodeBoard() string {
	var sb strings.Builder
	for _, row := range board.Pieces {
		for _, piece := range row {
			if piece.isEmpty {
				sb.WriteByte('.')
			} else {
				sb.WriteString(piece.color.GetColorString())
			}
			if piece.isQueen {
				sb.WriteByte('q')
			}
		}
		sb.WriteByte(' ')
	}
	return sb.String()
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
