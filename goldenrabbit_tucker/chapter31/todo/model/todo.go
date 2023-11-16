package todo_model

import (
	"gt/chap31/ex/util"
	"time"
)

type (
	TodoInfo struct {
		Id         string `json:"id"`
		Content    string `json:"content"`
		Check      bool   `json:"check"`
		Writer     string `json:"writer"`
		WriteUnix  int64  `json:"writeUnix"`
		UpdateUnix int64  `json:"updateUnix"`
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
