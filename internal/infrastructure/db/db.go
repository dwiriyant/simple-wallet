package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = runMigrations(sqlDB)
	if err != nil {
		panic(err)
	}

	return db
}

func runMigrations(db *sql.DB) error {
	// Specify the path to the migration files
	migrationDir := "internal/infrastructure/db/migrations"

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	// Run goose to apply all available migrations
	if err := goose.Up(db, migrationDir); err != nil {
		return err
	}

	return nil
}
