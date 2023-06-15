package redisWrapper

import (
	"context"
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

func AddSortedSet(ctx context.Context, key string, score float64, member interface{}) error {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := rdb.ZAdd(ctx, key, z).Result()
	return err
}
