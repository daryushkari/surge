package domain

import "gorm.io/gorm"

// RequestThreshold has price coefficient for when requests pass predefined threshold
// PriceCoefficient and RequestThreshold columns should be unique cause request threshold
// should not have more than one coefficient and also coefficient must be unique and incremental
type RequestThreshold struct {
	gorm.Model
	PriceCoefficient float64 `gorm:"unique"`
	RequestThreshold int64   `gorm:"unique"`
}
