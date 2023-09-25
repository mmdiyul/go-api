package main

import (
	"github.com/joho/godotenv"
	"go-api/config/database"
	"log/slog"
	"runtime/debug"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		debug.PrintStack()
		return
	}

	database.InitDatabase()
}
