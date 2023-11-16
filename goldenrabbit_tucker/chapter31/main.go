package main

import (
	core_redis "gt/chap31/ex/core/cache/redis"
	core_server "gt/chap31/ex/core/server"
)

func main() {
	redisOptions := new(core_redis.RedisConnInitOptions)
	redisOptions.Host = "localhost"
	if err := core_redis.RedisPoolInit(redisOptions); err != nil {
		panic(err)
	}

	if err := core_server.EchoServerStart(":3250"); err != nil {
		panic(err)
	}
}
