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

func AddSortedSet(ctx context.Context, key string, score float64, member interface{}) error {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := rdb.ZAdd(ctx, key, z).Result()
	return err
}

func GetCount(ctx context.Context, key string, min string, max string) (int64, error) {
	result := rdb.ZCount(ctx, key, min, max)
	return result.Result()
}

func RemoveOldElements(ctx context.Context, key string, max string) error {
	result := rdb.ZRemRangeByScore(ctx, key, "0", max)
	_, err := result.Result()
	return err
}
