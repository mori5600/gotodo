package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/logging"
	"github.com/mori5600/gotodo/todo"
)

func main() {
	logger := logging.GetLogger()

	t := todo.NewTodo(1, "Learn Go")
	logger.Info(t.String())
}
