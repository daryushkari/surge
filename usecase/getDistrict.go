package usecase

import (
	"context"
	"fmt"
	"surge/config"
	"surge/pkg/getPolygons"
	redisWrapper "surge/pkg/redis"
	"time"
)

func GetDistrict(ctx context.Context, longitude, latitude float64) (districtId string, err error) {
	tehranList, err := GetTehranDistrictPolygons(ctx)
	return "12", nil
}

// GetTehranDistrictPolygons set open street data in cache ,so it doesn't need to call
// open street APIs for every request
func GetTehranDistrictPolygons(ctx context.Context) (*getPolygons.TehranDistrictList, error) {
	tehranList := &getPolygons.TehranDistrictList{}
	key := fmt.Sprintf("%s:tehranPolygons", config.GetCnf().ServiceName)
	if err := redisWrapper.Get(ctx, key, tehranList); err == nil {
		return tehranList, nil
	}

	tehranList, err := getPolygons.ReturnPolygons()
	if err != nil {
		return nil, err
	}

	go redisWrapper.Set(context.Background(), key, tehranList, time.Hour*24*30)
	return tehranList, nil
}
