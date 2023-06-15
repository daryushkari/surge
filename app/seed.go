package app

import (
	"gorm.io/gorm"
	"log"
	"surge/entity/domain"
)

// Seed database tables when run the app and log errors when some row fails to be inserted
func Seed(db *gorm.DB) {
	SeedReqThreshold(db)
}

func SeedReqThreshold(db *gorm.DB) {
	var reqThresholds = []domain.RequestThreshold{
		{PriceCoefficient: 1.1, RequestThreshold: 10000},
		{PriceCoefficient: 1.2, RequestThreshold: 20000},
	}

	for _, v := range reqThresholds {
		result := db.Create(&v)
		if result.Error != nil {
			log.Println("error occurred in seed:", result.Error, "reqThreshold:", v)
		}
	}
}
