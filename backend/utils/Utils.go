package utils

func GetTileColor(row int, col int) string {
	if (row+col)%2 == 0 {
		return "w"
	} else {
		return "b"
	}
}

func IsInBounds(boardSize int, row int, col int) bool {
	return (row >= 0 && row < boardSize &&
		col >= 0 && col < boardSize)
}
