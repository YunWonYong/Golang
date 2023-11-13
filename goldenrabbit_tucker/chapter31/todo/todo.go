package todo

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	todo struct {
		Id    string `json:"id"`
		Content  string `json:"content"`
		Check   bool  `json:"check"`
		writeUnix int64
	}

	todoManager struct {
		info map[string]*todo
	}
)

var Manager *todoManager

func NewTodo() *todo {
	return new(todo)
}

func init() {
	Manager = new(todoManager)
	Manager.info = make(map[string]*todo)
	uuidV4 := Manager.genId()
	nowUnix := time.Now().Unix()
	Manager.info[uuidV4] = &todo{
		Id:    uuidV4,
		Content:  "Todo1",
		writeUnix: nowUnix,
	}

	uuidV4 = Manager.genId()
	Manager.info[uuidV4] = &todo{
		Id:    uuidV4,
		Content:  "Todo2",
		writeUnix: nowUnix + 3,
	}
}

func (m *todoManager) genId() string {
	for {
		id := uuid.NewString()
		_, exists := m.info[id]
		if exists {
			continue
		}
		return id
	}
}

func (m *todoManager) FindAll() []todo {
	list := make([]todo, 0, len(m.info))
	for _, td := range m.info {
		list = append(list, *td)
	}

	sort.Slice(list, func(prev, next int) bool {
		return list[prev].writeUnix < list[next].writeUnix
	})
	return list
}

func (m *todoManager) Find(id string) (todo, error) {
	td, exists := m.info[id]
	if exists {
		return *td, nil
	}

	return todo{}, errors.Errorf("todo id: %s not found ", id)
}

func (m *todoManager) Insert(td *todo) *todo {
	id := m.genId()
	td.Id = id
	td.writeUnix = time.Now().Unix()
	m.info[id] = td
	return td
}

func (m *todoManager) Update(id string) error {
	oldtd, err := m.Find(id)
	if err != nil {
		return err
	}
	
	oldtd.Check = !oldtd.Check

	m.info[id] = &oldtd
	return nil
}

func (m *todoManager) Delete(id string) error {
	delete(m.info, id)
	return nil
}
