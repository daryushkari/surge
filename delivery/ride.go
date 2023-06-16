package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"surge/entity/requestModel"
	"surge/pkg/errorMsg"
	"surge/usecase"
)

func SaveRide(ctx *gin.Context) {
	var input requestModel.SaveRideRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg.BadRequest})
		return
	}

	resp, err := usecase.SaveRide(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMsg.InternalServerError,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
