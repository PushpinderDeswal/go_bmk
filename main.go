package main

import (
	"github.com/PushpinderDeswal/go_bmk/cmd"
	"github.com/PushpinderDeswal/go_bmk/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.GetDatabase()

	defer database.CloseDatabase()
	cmd.SetDatabase(db)

	cmd.Execute()
}
