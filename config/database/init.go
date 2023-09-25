package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"runtime/debug"
)

func InitDatabase() *gorm.DB {
	dbDriver := os.Getenv("db.driver")
	if dbDriver != "pgsql" {
		slog.Error("Invalid database driver")
		return nil
	}

	dbHost := os.Getenv("db.host")
	dbPort := os.Getenv("db.port")
	dbName := os.Getenv("db.name")
	dbUsername := os.Getenv("db.username")
	dbPassword := os.Getenv("db.password")
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		debug.PrintStack()
	} else {
		slog.Info("Database connected!")
	}

	Migrate(db)

	return db
}
