package todo_model

import (
	"gt/chap31/ex/util"
	"time"
)

type (
	TodoInfo struct {
		Id        string `json:"id"`
		Content   string `json:"content"`
		Check     bool   `json:"check"`
		Writer    string
		WriteUnix int64
	}
)

func NewTodoInfo(content, writer string) *TodoInfo {
	todoInfo := new(TodoInfo)
	todoInfo.Id = util.GenUUID()
	todoInfo.Writer = writer
	todoInfo.Content = content
	todoInfo.WriteUnix = time.Now().Unix()
	return todoInfo
}
