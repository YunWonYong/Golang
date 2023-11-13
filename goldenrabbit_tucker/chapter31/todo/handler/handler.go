package todo_handler

import (
	"encoding/json"
	"gt/chap31/ex/todo"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func List(response http.ResponseWriter, request *http.Request) {
	todoList := todo.Manager.FindAll()
	buff, err := json.Marshal(todoList)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
		
	response.WriteHeader(http.StatusOK)
	response.Write(buff)
}

func Regist(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body) 
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	t :=  todo.NewTodo()
	if err := json.Unmarshal(body, t); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	t = todo.Manager.Insert(t)
	buff, err := json.Marshal(t)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(buff)
}

func Modify(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if err := todo.Manager.Update(id); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("ok"))
}

func Destroy(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if err := todo.Manager.Delete(id); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("ok"))
}
