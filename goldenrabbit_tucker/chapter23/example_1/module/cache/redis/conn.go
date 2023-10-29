package module_cache

import (
	"fmt"
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

func (conn *redisConn) do(command string, args ...interface{}) ([]byte, error) {
	reply, err := conn.Do(command, args...)
	if err != nil {
		return nil, err
	}

	var buff []byte
	switch t := reply.(type) {
	case string:
		buff = []byte(reply.(string))
	case []byte:
		buff = reply.([]byte)
	default:
		return nil, fmt.Errorf("reply to []byte fail. reply: %#+v, type: %#+v", reply, t)

	}
	return buff, nil
}
