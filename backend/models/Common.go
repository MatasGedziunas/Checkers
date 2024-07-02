package models

import "log"

type Coordinates struct {
	Row int
	Col int
}

func NewCoordinates(row, col int) Coordinates {
	return Coordinates{
		Row: row,
		Col: col,
	}
}

type Piece interface {
	GetCoordinates() Coordinates
	GetPossibleMoves() []Coordinates
}

type Color string

const (
	White Color = "w"
	Black Color = "b"
)

func (c Color) GetColorString() string {
	if c == White {
		return "w"
	} else if c == Black {
		return "b"
	} else {
		log.Fatalf("Invalid Color given: %v ; Expected White or Black of type Color", c)
		return ""
	}
}
