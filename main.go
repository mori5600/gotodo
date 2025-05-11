package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
)

func main() {
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
			name TEXT NOT NULL,
			email TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// データの挿入
	_, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "山田太郎", "yamada@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// データの取得
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 結果の表示
	for rows.Next() {
		var id int
		var name, email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, 名前: %s, メール: %s\n", id, name, email)
	}
}
