package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
	"github.com/mori5600/gotodo/logging"
)

func initDB() {
	logger := logging.GetLogger()

	// データベースに接続
	db, err := sql.Open(common.DB_DRIVER, common.DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルの作成
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		status INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		deleted_at DATETIME DEFAULT NULL
	)
`)
	if err != nil {
		logger.Error("Failed to create table", "error", err)
		return
	}

	logger.Info("Table created successfully")
}

func main() {
	initDB()
}
