package main

import (
	"fmt"
	"net/http"
	"strings"
)

type (
	handlerFunc      func(r *http.Request) (int, []byte, error)
	handlerTableType map[string]handlerFunc
)

var handlerTable = make(handlerTableType)

func (ht handlerTableType) Regist(path string, handler handlerFunc) {

	handlerKey := ht.GetPathByHandlerKey(path)

	if ht.DuplicatePathCheck(handlerKey) {
		panic(fmt.Sprintf("duplicate handler path: %s", path))
	}

	ht[handlerKey] = handler
}

func (ht handlerTableType) GetPathByHandlerKey(path string) string {
	handlerKey := path
	if ht.PathParamCheck(path) {
		handlerKey = ht.GetPathParamHandlerKey(path)
	}
	return handlerKey
}

func (ht handlerTableType) DuplicatePathCheck(handlerKey string) bool {
	_, exists := ht[handlerKey]
	return exists
}

func (ht handlerTableType) PathParamCheck(path string) bool {
	return strings.Contains(path, ":")
}

func (ht handlerTableType) GetPathParamHandlerKey(path string) string {
	pathParams := strings.Split(path, ":")
	return fmt.Sprintf("%s:%d", pathParams[0], len(pathParams)-1)
}

func (ht handlerTableType) GetPathByHandlerFunc(path string) handlerFunc {
	handlerFunc, exists := ht[path]
	if exists {
		return handlerFunc
	}

	paths := strings.Split(path, "/")
	temp := make([]string, 0)
	start := len(paths) - 1

	for start > -1 {
		temp = append(temp, paths[start])
		handlerKey := strings.Join(paths[0:start], "/")
		pathParamHandlerKey := fmt.Sprintf("%s/:%d", handlerKey, len(temp))
		hander, exists := ht[pathParamHandlerKey]
		if exists {
			return hander
		}
		start--
	}

	return func(r *http.Request) (int, []byte, error) {
		return http.StatusNotFound, []byte(fmt.Sprintf("not found handler. path: %s", path)), nil
	}
}

func main() {

	handlerTable.Regist("/", rootHandler)
	handlerTable.Regist("/users", usersHandler)
	handlerTable.Regist("/users/:id", usersPathParamsHandler)
	handlerTable.Regist("/users/:id/:name", usersPathParamsHandler2)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode, responseBuff := routing(r)
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBuff))
	})

	http.ListenAndServe(":3000", nil)
}

func routing(r *http.Request) (statusCode int, responseBuff []byte) {
	path := r.URL.Path
	handler := handlerTable.GetPathByHandlerFunc(path)
	var err error = nil
	statusCode, responseBuff, err = handler(r)
	if err != nil {
		statusCode = http.StatusInternalServerError
		responseBuff = []byte("server error")
	}
	return
}

func rootHandler(r *http.Request) (int, []byte, error) {
	return http.StatusOK, []byte("hello world!!!"), nil
}

func usersHandler(r *http.Request) (int, []byte, error) {
	return http.StatusOK, []byte("users"), nil
}

func usersPathParamsHandler(r *http.Request) (int, []byte, error) {
	paths := strings.Split(r.URL.Path, "/")
	return http.StatusOK, []byte(fmt.Sprintf("path: %s, pathsParams: %s", r.URL.Path, strings.Join(paths[2:], ","))), nil
}

func usersPathParamsHandler2(r *http.Request) (int, []byte, error) {
	paths := strings.Split(r.URL.Path, "/")
	return http.StatusOK, []byte(fmt.Sprintf("path: %s, pathsParams: %s", r.URL.Path, strings.Join(paths[2:], ","))), nil
}
