package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surge/entity/requestModel"
)

func SaveRide(ctx *gin.Context) {
	var input requestModel.SaveRideRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestModel.BadRequestError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": requestModel.SaveRideResponse{
			Message: requestModel.SuccessfulMessage,
			Code:    http.StatusOK,
		},
	})
}
