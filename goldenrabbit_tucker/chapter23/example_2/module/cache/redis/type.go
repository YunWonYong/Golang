package module_cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type (
	RedisConnInitOptions struct {
		ConnectLimitFlag   bool
		Network            string
		Host               string
		address            string
		Port               int64
		ConnectLimitSecond time.Duration
		PoolSize           int
		Slot               int
	}
)

type (
	redisConn struct {
		redis.Conn
	}

	redisConnPool struct {
		*redis.Pool
		options *RedisConnInitOptions
	}
)
