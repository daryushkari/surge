package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surge/entity/requestModel"
	"surge/pkg/errorMsg"
)

func SaveRide(ctx *gin.Context) {
	var input requestModel.SaveRideRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg.BadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": requestModel.SaveRideResponse{
			Message: errorMsg.BadRequest,
			Code:    http.StatusOK,
		},
	})
}
