package domain

import "gorm.io/gorm"

type RequestThreshold struct {
	gorm.Model
	PriceCoefficient float64 `gorm:"unique"`
	RequestThreshold int64   `gorm:"unique"`
}
