package core_redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var connPool *redisConnPool

func RedisPoolInit(options *RedisConnInitOptions) error {
	pool := new(redis.Pool)
	pool.Dial = getRedisDial(options)
	pool.MaxIdle = options.PoolSize

	connPool = new(redisConnPool)
	connPool.options = options
	return connPool.init(pool)
}

func (rcp *redisConnPool) init(pool *redis.Pool) (err error) {
	connPool.Pool = pool
	poolSize := rcp.options.PoolSize
	connList := make([]*redisConn, 0, poolSize)
	for i := 0; i < poolSize; i++ {
		conn := GetConn()
		connList = append(connList, conn)
		if err = conn.Ping("redis pool ping test"); err != nil {
			err = fmt.Errorf("index: %d, conn ping test fail. err: %s", i, err.Error())
			break
		}

	}
	return connPool.initClean(connList, err)
}

func (rcp *redisConnPool) initClean(connList []*redisConn, err error) error {
	for i, conn := range connList {
		if closeErr := conn.Close(); closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%s\nindex: %d, conn close fail. error: %s", err.Error(), i, closeErr.Error())
				continue
			}
			err = fmt.Errorf("index: %d, conn close fail. error: %s", i, closeErr.Error())
		}
	}
	return err
}

func GetConn() *redisConn {
	conn := connPool.Get()
	rc := new(redisConn)
	rc.Conn = conn
	return rc
}
