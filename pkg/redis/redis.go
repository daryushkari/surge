package redis

import (
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitClient(addr string) {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr: addr,
		})
	}
}

func GetClient() *redis.Client {
	return rdb
}
