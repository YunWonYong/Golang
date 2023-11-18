package core_redis

import (
	"gt/chap31/ex/util"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
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

type (
	commandByResult interface {
		GetByte(input interface{}) (output []byte, err error)
	}

	getCommandResult []byte

	setCommandResult struct {
		AffectedCount int64 `json:"affectedCount"`
	}

	keyValueCommandResult struct {
		Keys   []string `json:"keys"`
		Values []string `json:"values"`
	}

	mapCommandResult struct {
		*keyValueCommandResult
		Info map[string]string `json:"info"`
	}
)

func (gc *getCommandResult) GetByte(input interface{}) (output []byte, err error) {
	return util.InterfaceToByteArray(input)
}

func (sc *setCommandResult) GetByte(input interface{}) (output []byte, err error) {
	number, err := util.InterfaceToNumber[int64](input)
	if err != nil {
		return nil, err
	}

	sc.AffectedCount = number

	return util.Marshal(sc)
}

func (kvc *keyValueCommandResult) GetByte(input interface{}) (output []byte, err error) {
	arr, ok := input.([]interface{})
	if !ok {
		err = errors.New("input data type not []interface{}")
		return
	}

	index := 0
	size := len(arr)
	kvc.Keys = make([]string, 0, size)
	kvc.Values = make([]string, 0, size)
	for index < size {
		key, err := util.InterfaceToString(arr[index])
		if err != nil {
			return nil, errors.WithMessagef(err, "key: %#v", arr[index])
		}

		value, err := util.InterfaceToString(arr[index+1])

		if err != nil {
			return nil, errors.WithMessagef(err, "value: %#v", arr[index+1])
		}
		kvc.Keys = append(kvc.Keys, key)
		kvc.Values = append(kvc.Values, value)
		index += 2
	}

	return util.Marshal(kvc)
}

func (mc *mapCommandResult) GetByte(input interface{}) (output []byte, err error) {
	mc.keyValueCommandResult = new(keyValueCommandResult)
	_, err = mc.keyValueCommandResult.GetByte(input)
	if err != nil {
		return
	}

	mc.Info = make(map[string]string)
	for index, key := range mc.keyValueCommandResult.Keys {
		mc.Info[key] = mc.keyValueCommandResult.Values[index]
	}

	return util.Marshal(mc)
}
