package usecase

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"surge/config"
	"surge/entity/requestModel"
	"surge/pkg/errorMsg"
	redisWrapper "surge/pkg/redis"
	"surge/repository"
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

	min := fmt.Sprintf("%d", time.Now().UnixMilli()-config.GetCnf().RequestTimeWindow)
	max := fmt.Sprintf("%d", time.Now().UnixMilli())
	res, err := redisWrapper.GetCount(ctx, key, min, max)
	if err != nil {
		return requestModel.SaveRideResponse{
			Message: errorMsg.InternalServerError,
			Code:    http.StatusInternalServerError,
		}, err
	}

	coe, err := GetRequestCoefficient(ctx, res)
	if err != nil {
		return requestModel.SaveRideResponse{
			Message: errorMsg.InternalServerError,
			Code:    http.StatusInternalServerError,
		}, err
	}

	go NotifyPricing(ctx, coe, disId)

	return requestModel.SaveRideResponse{
		Message: errorMsg.SuccessfullySaved,
		Code:    http.StatusOK,
	}, nil
}

// GetRequestCoefficient returns 1 when requests are lower than predefined thresholds
func GetRequestCoefficient(ctx context.Context, requestCount int64) (coe float64, err error) {
	thresholdList, err := repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	coe = 1
	for _, v := range thresholdList {
		if requestCount > v.RequestThreshold {
			coe = v.PriceCoefficient
		}
	}
	return coe, nil
}
