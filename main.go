package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api/app/routes"
	"go-api/config/database"
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		slog.Error(errEnv.Error())
		debug.PrintStack()
		return
	}

	r := gin.Default()
	db := database.InitDatabase()

	routes.Register(r, db)

	errRun := r.Run(os.Getenv("app.url"))
	if errRun != nil {
		slog.Error(errRun.Error())
		debug.PrintStack()
		return
	}
}
