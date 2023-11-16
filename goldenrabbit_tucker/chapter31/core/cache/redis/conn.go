package core_redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

const (
	SET     string = "SET"
	GET     string = "GET"
	PING    string = "PING"
	HSET    string = "HSET"
	HDEL    string = "HDEL"
	HGET    string = "HGET"
	HGETALL string = "HGETALL"
)

func (conn *redisConn) Ping(pingMsg string) error {
	_, err := conn.do("PING", pingMsg)
	return err
}

func (conn *redisConn) Close() error {
	if err := conn.Conn.Close(); err != nil {
		err = fmt.Errorf("redisConn Close fail. err: %s", err.Error())
		return err
	}
	return nil
}

func (conn *redisConn) HSET(key string, member, buff interface{}) error {
	_, err := conn.do(HSET, key, member, buff)
	if err != nil {
		err = errors.WithMessagef(err, "key: %s, member: %#v, buff: %#v", key, member, buff)
	}
	return err
}

func (conn *redisConn) HGET(key string, member interface{}) (buff []byte, err error) {
	buff, err = conn.do(HGET, key, member)
	if err != nil {
		err = errors.WithMessagef(err, "key: %s, member: %#v", key, member)
	}
	return
}

func (conn *redisConn) HDEL(key string, member interface{}) error {
	_, err := conn.do(HDEL, key, member)
	if err != nil {
		err = errors.WithMessagef(err, "key: %s, member: %#v", key, member)
	}
	return err
}

func (conn *redisConn) HGETALL(key string) (buff []byte, err error) {
	buff, err = conn.do(HGETALL, key)
	if err != nil {
		err = errors.WithMessagef(err, "key: %s", key)
	}
	return
}

func (conn *redisConn) do(command string, args ...interface{}) (buff []byte, err error) {
	reply, err := conn.Do(command, args...)
	if err == redis.ErrNil {
		return []byte{}, nil
	} else if err != nil {
		return nil, err
	} else if reply == nil {
		return []byte{}, errors.Errorf("result nil.")
	}

	buff, err = CommandByByte(command, reply)
	if err == redis.ErrNil {
		return []byte{}, nil
	}
	return buff, err
}

func CommandByByte(command string, obj interface{}) (buff []byte, err error) {
	var cbr commandByResult
	switch command {
	case PING:
		fallthrough
	case HGET:
		fallthrough
	case GET:
		cbr = new(getCommandResult)
	case HGETALL:
		cbr = new(mapCommandResult)
	case HSET:
		fallthrough
	case HDEL:
		fallthrough
	case SET:
		cbr = new(setCommandResult)
	default:
		return nil, errors.Errorf("not supported command(%s).", command)
	}
	return cbr.GetByte(obj)
}
