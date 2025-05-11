package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("アプリケーションが起動しました", "version", "1.0.0", "port", 8080)
}
