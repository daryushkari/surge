package redisWrapper

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

const (
	NoTTLTime = -1
)

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

func Get(ctx context.Context, key string, dest interface{}) error {
	p, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(p), &dest)
}

func Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, key, p, duration).Err()
}
