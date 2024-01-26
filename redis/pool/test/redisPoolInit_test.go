package pool_test

import (
	"context"
	"testing"
	"time"

	"github.com/YunWonYong/redis/pool"
)

func TestRedisPoolInit(t *testing.T) {
	options := new(pool.RedisPoolInitOptions)
	options.Network = "tcp"
	options.Address = "localhost:6379"
	timeoutSec := time.Duration(5) * time.Second
	options.ConnectTimeout = timeoutSec
	options.WriteTimeout = timeoutSec
	options.ReadTimeout = timeoutSec
	options.MaxIdle = 10
	options.Database = 0

	p, err := pool.NewRedisPool(context.TODO(), options)

	if err != nil {
		t.Error(err)
		return
	}

	if p == nil {
		t.Fail()
	}

	if err := pool.PingTest(p.Pool); err != nil {
		t.Error(err)
		return
	}
}

func TestRedisPoolInitPingTest(t *testing.T) {
	options := new(pool.RedisPoolInitOptions)
	options.Network = "tcp"
	options.Address = "localhost:6379"
	timeoutSec := time.Duration(5) * time.Second
	options.ConnectTimeout = timeoutSec
	options.WriteTimeout = timeoutSec
	options.ReadTimeout = timeoutSec
	options.MaxIdle = 10
	options.Database = 0
	options.PingTestFlag = true

	if _, err := pool.NewRedisPool(context.TODO(), options); err != nil {
		t.Error(err)
		return
	}
}
