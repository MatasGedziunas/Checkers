package utils

func GetTileColor(row int, col int) string {
	if (row+col)%2 == 0 {
		return "w"
	} else {
		return "b"
	}
}

func IsInBounds(row int, col int, boardSize int) bool {
	return (row >= 0 && row < boardSize &&
		col >= 0 && col < boardSize)
}

func Remove[T any](slice []T, indexToRemove int) []T {
	return append(slice[:indexToRemove], slice[indexToRemove+1:]...)
}

func GetDirection(fromCell int, toCell int) int {
	if fromCell > toCell {
		return -1
	} else {
		return 1
	}
}
