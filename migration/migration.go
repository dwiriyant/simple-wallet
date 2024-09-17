package migration

import (
	"log"
	"simple-wallet/db"
	"simple-wallet/models"
)

func Migrate() {
	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Wallet{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
