package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
	"github.com/mori5600/gotodo/logging"
)

func main() {
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
		logger.Error("テーブルの作成に失敗しました", "error", err)
	}

	logger.Info("テーブルの作成に成功しました")
}
