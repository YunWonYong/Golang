package main

import (
	"fmt"
	module_cache "gt/chap23/ex2/ex/module/cache/redis"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("main thread panic.\n %+v\n", r)
		}
	}()

	if err := redisPoolInit(); err != nil {
		panic(fmt.Errorf("redisPool Init fail.\n err: %s", err.Error()))
	}
}

func redisPoolInit() error {
	options := new(module_cache.RedisConnInitOptions)
	options.Host = "localhost"
	options.Slot = 0
	options.PoolSize = 20
	options.ConnectLimitFlag = true
	return module_cache.RedisPoolInit(options)
}
