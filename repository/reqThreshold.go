package repository

import (
	"context"
	"fmt"
	"surge/config"
	"surge/entity/domain"
	postgresql "surge/pkg/postgis"
	redisWrapper "surge/pkg/redis"
	"time"
)

const (
	RequestThresholdKey = "requestThresholds"
)

// GetAll read from redis cache and if it's empty read configs from database
func GetAll(ctx context.Context) (thresholdList []domain.RequestThreshold, err error) {
	key := fmt.Sprintf("%s:%s", config.GetCnf().ServiceName, RequestThresholdKey)
	if err = redisWrapper.Get(ctx, key, &thresholdList); err == nil {
		return thresholdList, nil
	}

	result := postgresql.Get().Find(&thresholdList)
	if result.Error != nil {
		return nil, result.Error
	}
	go redisWrapper.Set(context.Background(), key, thresholdList, time.Hour*24)
	return thresholdList, nil
}
