package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
)

func InitDB() (*sql.DB, error) {
	// Create a new SQLite database

	db, err := sql.Open(common.DB_DRIVER, common.DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// If the database file does not exist, it will be created
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		status INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		deleted_at DATETIME DEFAULT NULL
	)
`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
