package databaseFunctionality

// all database models should have first field as id

type Game struct {
	GameId      int `json:"game_id" db:"game_id"`
	WhitePlayer int `json:"white_player" db:"white_player"`
	BlackPlayer int `json:"black_player" db:"black_player"`
}

func (g Game) Validate() (string, bool) {
	if g.BlackPlayer == g.WhitePlayer {
		return "black_player id cannot equal white_player_id", false
	}
	return "", true
}

type User struct {
	UserId   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (u User) Validate() (string, bool) {
	return "", true
}

type Board struct {
	BoardId     int    `json:"board_id" db:"board_id"`
	BoardString string `json:"board_string" db:"board_string"`
	GameId      int    `json:"game_id" db:"game_id"`
}

func (b Board) Validate() (string, bool) {
	return "", true
}
