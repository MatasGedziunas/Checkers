package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func createTables(db *sql.DB) error {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        user_id SERIAL PRIMARY KEY,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("could not create users table: %v", err)
	}

	createGamesTable := `
    CREATE TABLE IF NOT EXISTS games (
        game_id SERIAL PRIMARY KEY,
        white_player INTEGER NOT NULL,
        black_player INTEGER NOT NULL,
		FOREIGN KEY (white_player) REFERENCES users(user_id) ON DELETE CASCADE,
		FOREIGN KEY (black_player) REFERENCES users(user_id) ON DELETE CASCADE
    );`
	_, err = db.Exec(createGamesTable)
	if err != nil {
		return fmt.Errorf("could not create games table: %v", err)
	}

	createBoardsTable := `
    CREATE TABLE IF NOT EXISTS boards (
        board_id SERIAL PRIMARY KEY,
        board_string TEXT NOT NULL,
        game_id INTEGER NOT NULL,
        FOREIGN KEY (game_id) REFERENCES games(game_id) ON DELETE CASCADE
    );`
	_, err = db.Exec(createBoardsTable)
	if err != nil {
		return fmt.Errorf("could not create boards table: %v", err)
	}
	return nil
}

func dropTables(db *sql.DB) error {
	dropBoardsTable := `DROP TABLE IF EXISTS boards;`
	_, err := db.Exec(dropBoardsTable)
	if err != nil {
		return fmt.Errorf("could not drop boards table: %v", err)
	}

	dropGamesTable := `DROP TABLE IF EXISTS games;`
	_, err = db.Exec(dropGamesTable)
	if err != nil {
		return fmt.Errorf("could not drop games table: %v", err)
	}

	dropUsersTable := `DROP TABLE IF EXISTS users;`
	_, err = db.Exec(dropUsersTable)
	if err != nil {
		return fmt.Errorf("could not drop users table: %v", err)
	}

	return nil
}

func connectDB() (*sql.DB, error) {
	password := os.Getenv("postgresql_password")
	fmt.Println(password)
	fmt.Println("hello")
	connStr := fmt.Sprintf("user=postgres dbname=checkers sslmode=disable password=%s", password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	err = dropTables(db)
	if err != nil {
		fmt.Println(err)
	}
	err = createTables(db)
	if err != nil {
		fmt.Println(err)
	}
}
