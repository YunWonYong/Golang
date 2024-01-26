package pool

import (
	"context"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

type (
	RedisPoolInitOptions struct {
		Network        string
		Address        string
		Database       int
		ConnectTimeout time.Duration
		WriteTimeout   time.Duration
		ReadTimeout    time.Duration
		MaxIdle        int
		PingTestFlag   bool
	}

	redisPool struct {
		*redis.Pool
	}
)

func (rp *redisPool) Get() redis.Conn {
	return rp.Get()
}

func NewRedisPool(ctx context.Context, opts *RedisPoolInitOptions) (*redisPool, error) {
	p := new(redisPool)
	p.Pool = &redis.Pool{Dial: getDial(opts)}

	if opts.MaxIdle > 0 {
		p.Pool.MaxIdle = opts.MaxIdle
	}

	if opts.PingTestFlag {
		if err := PingTest(p.Pool); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func PingTest(pool *redis.Pool) error {
	var (
		conns []redis.Conn = make([]redis.Conn, 0)
		conn  redis.Conn
		i     int
		err   error
	)

	for i < pool.MaxIdle {
		conn = pool.Get()
		if err = conn.Err(); err != nil {
			return err
		}

		if _, err = conn.Do("PING"); err != nil {
			break
		}

		conns = append(conns, conn)
		i++
	}

	if err == nil {
		return pingTestConnsClose(conns)
	}

	return nil
}

func getDial(opts *RedisPoolInitOptions) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {
		return redis.Dial(
			opts.Network,
			opts.Address,
			redis.DialReadTimeout(opts.ReadTimeout),
			redis.DialWriteTimeout(opts.WriteTimeout),
			redis.DialDatabase(opts.Database),
		)
	}
}

func pingTestConnsClose(conns []redis.Conn) error {
	for _, conn := range conns {
		if err := conn.Close(); err != nil {
			return err
		}
	}

	return nil
}
