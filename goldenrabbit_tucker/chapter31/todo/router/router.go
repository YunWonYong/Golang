package todo_router

import (
	todo_handler "gt/chap31/ex/todo/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Load(router *mux.Router) {
	router.HandleFunc("/todo", todo_handler.List).Methods(http.MethodGet)
	router.HandleFunc("/todo", todo_handler.Regist).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", todo_handler.Modify).Methods(http.MethodPut)
	router.HandleFunc("/todo/{id}", todo_handler.Destroy).Methods(http.MethodDelete)
}
