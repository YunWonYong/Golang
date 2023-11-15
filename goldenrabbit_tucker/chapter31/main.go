package main

import (
	core_redis "gt/chap31/ex/core/cache/redis"
	core_server "gt/chap31/ex/core/server"
)

func main() {
	redisOptions := new(core_redis.RedisConnInitOptions)
	redisOptions.Host = "localhost"
	core_redis.RedisPoolInit(redisOptions)
	if err := core_server.EchoServerStart(":3250"); err != nil {
		panic(err)
	}
}
