package todo_dao

import (
	"fmt"
	core_redis "gt/chap31/ex/core/cache/redis"
	todo_model "gt/chap31/ex/todo/model"
	"gt/chap31/ex/util"
	"time"

	"github.com/pkg/errors"
)

const (
	TODO_KEY = "TODO"
)

func List(writer string) ([]byte, error) {
	conn := core_redis.GetConn()
	defer conn.Close()

	key := getTodoKey(writer)
	buff, err := conn.HGETALL(key)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func Insert(dto *todo_model.TodoRegistDTO) error {
	insertInfo := todo_model.NewTodoInfo(dto.Content, dto.Writer)

	conn := core_redis.GetConn()
	defer conn.Close()
	key := getTodoKey(dto.Writer)

	buff, err := util.Marshal(insertInfo)

	if err != nil {
		return err
	}

	if err := conn.HSET(key, insertInfo.Id, buff); err != nil {
		return errors.WithMessage(err, "todo save fail.")
	}

	return nil
}

func Toggle(writer, id string) error {
	conn := core_redis.GetConn()
	defer conn.Close()

	key := getTodoKey(writer)
	buff, err := conn.HGET(key, id)
	if err != nil {
		return errors.WithMessagef(err, "writer: %s, id: %s not found todo data.", writer, id)
	}
	fmt.Printf("toggle old data: %s\n", buff)
	info, err := util.Unmarshal[todo_model.TodoInfo](buff)

	if err != nil {
		return err
	}

	info.Check = !info.Check
	info.UpdateUnix = time.Now().Unix()
	buff, err = util.Marshal(info)
	fmt.Printf("toggle new data: %s\n", buff)
	if err != nil {
		return err
	}

	return conn.HSET(key, id, buff)
}

func Delete(writer, id string) error {
	conn := core_redis.GetConn()
	defer conn.Close()

	key := getTodoKey(writer)

	return conn.HDEL(key, id)
}

func Truncate(writer string) error {
	conn := core_redis.GetConn()
	defer conn.Close()
	key := getTodoKey(writer)
	return conn.DEL(key)
}

func getTodoKey(writer string) string {
	return fmt.Sprintf("%s:%s", TODO_KEY, writer)
}
