package models

import (
	"log"
)

type Checker struct {
	cords   Coordinates
	isEmpty bool
	isQueen bool
	color   Color
}

func NewChecker(row int, col int, color string) Checker {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	var tempColor Color = getColorFromString(color)
	return Checker{cords: cords,
		color: tempColor,
	}
}

func NewEmptyChecker(row int, col int) Checker {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	return Checker{cords: cords,
		isEmpty: true}
}

func NewQueen(row int, col int, color string) Checker {
	cords := Coordinates{
		Row: row,
		Col: col,
	}
	var tempColor Color = getColorFromString(color)
	return Checker{cords: cords,
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

func (checker *Checker) GetCoordinates() Coordinates {
	return checker.cords
}

func (checker *Checker) GetPossibleMoves(board *Board) []Coordinates {
	if checker.isQueen {
		return checker.getQueenPossibleMoves(*board)
	} else {
		return checker.getCheckerPossibleMoves(*board)
	}
}

func (checker Checker) getCheckerPossibleMoves(board Board) []Coordinates {
	possibleMoves := []Coordinates{}
	return possibleMoves
}

func (queen Checker) getQueenPossibleMoves(board Board) []Coordinates {
	possibleMoves := []Coordinates{}
	return possibleMoves
}
