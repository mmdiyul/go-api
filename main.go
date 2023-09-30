package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go-api/app/routes"
	"go-api/config/database"
	"log/slog"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
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

	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	routes.Register(r, db)

	errRun := r.Run(os.Getenv("app.url"))
	if errRun != nil {
		slog.Error(errRun.Error())
		debug.PrintStack()
		return
	}
}
