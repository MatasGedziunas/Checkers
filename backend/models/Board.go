package models

import (
	"strings"
)

type Board struct {
	Pieces [][]Checker
}

func NewBoard(boardString string) Board {
	board := strings.Split(boardString, " ")
	var pieces [][]Checker
	var row []Checker
	curRow := 0
	for curRow < len(board)-1 {
		row = []Checker{}
		curStr := 0
		for curStr < len(board[curRow]) {
			var piece Checker
			if board[curRow][curStr] == '.' {
				piece = NewEmptyChecker(curRow, len(row))
			} else if len(board[curRow]) < curStr+1 && board[curRow][curStr+1] == 'q' {
				piece = NewQueen(curRow, len(row), string(board[curRow][curStr]))
				curStr += 1
			} else {
				piece = NewChecker(curRow, len(row), string(board[curRow][curStr]))
			}
			row = append(row, piece)
			curStr += 1
		}
		pieces = append(pieces, row)
		curRow += 1
	}
	return Board{pieces}
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
