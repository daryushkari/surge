package cronjob

import (
	"context"
	"fmt"
	"log"
	"surge/config"
	"surge/pkg/getPolygons"
	redisWrapper "surge/pkg/redis"
	"time"
)

func RemoveOldRequest() {
	cnf := config.GetCnf()
	tehranList := &getPolygons.TehranDistrictList{}
	key := fmt.Sprintf("%s:tehranPolygons", cnf.ServiceName)
	if err := redisWrapper.Get(context.Background(), key, tehranList); err != nil {
		log.Println("error getting lists", err)
	}

	for _, v := range tehranList.Districts {
		dKey := fmt.Sprintf("%s:%s", config.GetCnf().ServiceName, v)
		max := fmt.Sprintf("%d", time.Now().UnixMilli()-config.GetCnf().RequestLiveTime)
		err := redisWrapper.RemoveOldElements(context.Background(), dKey, max)
		if err != nil {
			log.Println("error removing old requests", err)
		}
	}
}
