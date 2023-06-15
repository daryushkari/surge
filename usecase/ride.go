package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"surge/config"
	"surge/entity/requestModel"
	"surge/pkg/errorMsg"
	redisWrapper "surge/pkg/redis"
	"time"
)

func SaveRide(ctx *gin.Context, req requestModel.SaveRideRequest) (requestModel.SaveRideResponse, error) {
	disId, err := GetDistrict(ctx, req.Longitude, req.Latitude)
	if err != nil {
		return requestModel.SaveRideResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}, err
	}

	nowTime := float64(time.Now().UnixMilli())
	key := fmt.Sprintf("%s:%s", config.GetCnf().ServiceName, disId)
	err = redisWrapper.AddSortedSet(ctx, key, nowTime, nowTime)
	if err != nil {
		return requestModel.SaveRideResponse{
			Message: errorMsg.InternalServerError,
			Code:    http.StatusInternalServerError,
		}, err
	}

	return requestModel.SaveRideResponse{
		Message: errorMsg.SuccessfullySaved,
		Code:    http.StatusOK,
	}, nil
}
