package migration

import (
	"log"

	"gorm.io/gorm"
)

// RunMigrations runs all database migrations using GORM
func DatabaseMigrations(db *gorm.DB) error {
	migration(db)
	return nil
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		// &entities.Users{},
	); err != nil {
		log.Println("🌟 GORM database migrations failed 🚀")
		return err
	}
	log.Println("🌟 GORM database migrations completed successfully 🚀")
	return nil
}