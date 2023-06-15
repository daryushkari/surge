package app

import (
	"gorm.io/gorm"
	"log"
	"surge/entity/domain"
)

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.RequestThreshold{})
	if err != nil {
		log.Fatalf("migration failed: %v", err.Error())
	}
	return nil
}
