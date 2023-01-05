package redisdemo

import (
	"sync"

	"github.com/go-redis/redis"
)

type IRedisConfig interface {
	GetPing() (string, error)
}

type RedisConfig struct {
	*redis.Client
}

var (
	redisClient *RedisConfig
	mtx         sync.Mutex
)

func GetRedisClient() *RedisConfig {
	mtx.Lock()
	defer mtx.Unlock()
	if redisClient == nil {
		redisClient = &RedisConfig{
			redis.NewClient(&redis.Options{
				Addr: "redis-server:6379",
			}),
		}
	}

	return redisClient
}

func (rc *RedisConfig) GetPing() (string, error) {
	return rc.Ping().Result()
}
