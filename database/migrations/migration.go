package migrations

import (
	"fmt"
	"log"

	"blog/database"
	"blog/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.Post{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
