package main

import module_cache "gt/chap23/ex1/ex/module/cache/redis"

func main() {
	if err := redisPoolInit(); err != nil {
		panic(err)
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
