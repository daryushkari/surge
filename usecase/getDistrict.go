package usecase

import (
	"github.com/gin-gonic/gin"
)

func GetDistrict(ctx *gin.Context, longitude, latitude float64) (districtId string, err error) {
	return "12", nil
}
