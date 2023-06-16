package usecase

import (
	"context"
	"errors"
	"fmt"
	"surge/config"
	"surge/pkg/getPolygons"
	redisWrapper "surge/pkg/redis"
	"time"
)

func GetDistrict(ctx context.Context, longitude, latitude float64) (districtId string, err error) {
	tehranList, err := GetTehranDistrictPolygons(ctx)

	for _, v := range tehranList.Polygons {
		if isPointInPolygon(longitude, latitude, v) {
			return v.DistrictId, nil
		}
	}
	return "", errors.New("coordinate not in tehran")
}

// isPointInPolygon checks if coordinate is in district with ray casting algorithm
func isPointInPolygon(longitude, latitude float64, district *getPolygons.DistrictPolygon) bool {
	isInside := false
	points := district.Points
	for i, j := 0, len(points)-1; i < len(points); i++ {
		if (points[i].Latitude < latitude && points[j].Latitude >= latitude) || (points[j].Latitude < latitude && points[i].Latitude >= latitude) {
			if points[i].Longitude+(latitude-points[i].Latitude)/(points[j].Latitude-points[i].Latitude)*(points[j].Longitude-points[i].Longitude) < longitude {
				isInside = !isInside
			}
		}
		j = i
	}
	return isInside
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
