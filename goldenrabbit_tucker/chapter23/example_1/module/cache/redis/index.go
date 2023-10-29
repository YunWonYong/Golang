package module_cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	REDIS_DEFAULT_OPTIONS_NETWORK           = "tcp"
	REDIS_DEFAULT_OPTIONS_HOST              = "localhost"
	REDIS_DEFAULT_OPTIONS_PORT              = 6379
	REDIS_DEFAULT_OPTIONS_CONN_LIMIT_SECOND = time.Duration(5) * time.Second
	REDIS_DEFAULT_OPTIONS_POOL_SZIE         = 10
)

func getRedisDial(options *RedisConnInitOptions) func() (redis.Conn, error) {
	setRedisDefaultOptions(options)
	return func() (redis.Conn, error) {
		dialOptions := make([]redis.DialOption, 0)
		if options.ConnectLimitFlag {
			dialOptions = append(
				dialOptions,
				redis.DialWriteTimeout(options.ConnectLimitSecond),
				redis.DialReadTimeout(options.ConnectLimitSecond),
				redis.DialConnectTimeout(options.ConnectLimitSecond),
			)
		}
		dialOptions = append(
			dialOptions,
			redis.DialDatabase(options.Slot),
		)
		return redis.Dial(
			options.Network,
			options.address,
			dialOptions...,
		)
	}
}

func setRedisDefaultOptions(options *RedisConnInitOptions) {
	if len(options.Network) == 0 {
		options.Network = REDIS_DEFAULT_OPTIONS_NETWORK
	}

	if len(options.Host) == 0 {
		options.Host = REDIS_DEFAULT_OPTIONS_HOST
	}

	if options.Port == 0 {
		options.Port = REDIS_DEFAULT_OPTIONS_PORT
	}

	if options.ConnectLimitFlag && options.ConnectLimitSecond == 0 {
		options.ConnectLimitSecond = REDIS_DEFAULT_OPTIONS_CONN_LIMIT_SECOND
	}

	if options.PoolSize == 0 {
		options.PoolSize = REDIS_DEFAULT_OPTIONS_POOL_SZIE
	}

	options.address = fmt.Sprintf("%s:%d", options.Host, options.Port)
}
