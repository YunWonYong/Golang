package student_router

import (
	student_handler "gt/chap30/ex/student/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Load(router *mux.Router) {
	router.HandleFunc("/student", student_handler.List).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}", student_handler.Read).Methods(http.MethodGet)
	router.HandleFunc("/student", student_handler.Regist).Methods(http.MethodPost)
	router.HandleFunc("/student/{id}", student_handler.Modify).Methods(http.MethodPut)
	router.HandleFunc("/student/{id}", student_handler.Destroy).Methods(http.MethodDelete)
}
