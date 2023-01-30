package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func NewRedis(opt RedisOpt) IRedis {
	i := RedisOpt{
		Address:  opt.Address,
		Password: opt.Password,
		Db:       opt.Db,
		Expired:  opt.Expired,
		ctx:      opt.ctx,
	}

	client := redis.NewClient(&redis.Options{
		Addr:     i.Address,
		Password: i.Password,
		DB:       i.Db,
	})

	rdb = client

	return i
}

func (i RedisOpt) Set(key string, val any) (bool, error) {
	return rdb.SetNX(key, val, i.Expired*time.Minute).Result()
}

func (i RedisOpt) Get(key string) (string, error) {
	return rdb.Get(key).Result()
}

func (i RedisOpt) Del(key string) (int64, error) {
	return rdb.Del(key).Result()
}
