package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type RedisOpt struct {
	Address  string
	Password string
	Db       int
	Expired  time.Duration
	ctx      context.Context
}

type IRedis interface {
	Set(key string, val any) (bool, error)
	Get(key string) (string, error)
	Del(key string) (int64, error)
}

var (
	REDIS_NIL      = redis.Nil
	KEY_NOT_EXISTS = errors.New("key not exists")
)
