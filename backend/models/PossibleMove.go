package models

type PossibleMove struct {
	CapturesCount int
	Move          Coordinates
}

func NewPossibleMove(row int, col int, capturesCount int) PossibleMove {
	return PossibleMove{
		CapturesCount: capturesCount,
		Move: Coordinates{
			Row: row,
			Col: col,
		},
	}
}
