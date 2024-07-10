package models

type PossibleMove struct {
	CapturesCount int
	Moves         []Coordinates
}

func NewPossibleMove(capturesCount int, moves []Coordinates) PossibleMove {
	return PossibleMove{
		CapturesCount: capturesCount,
		Moves:         moves,
	}
}
