package main

import (
	"fmt"
	"net/http"
	"strings"
)

type handlerFunc = func(r *http.Request) ([]byte, error)

var handlerTable = make(map[string]handlerFunc)

func main() {
	handlerTable["/"] = rootHandler
	handlerTable["/users"] = usersHandler

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode, responseBuff := routing(r)
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBuff))
	})

	http.ListenAndServe(":3000", nil)
}

func routing(r *http.Request) (statusCode int, responseBuff []byte) {
	statusCode = http.StatusInternalServerError
	responseBuff = []byte("server error")
	path := r.URL.Path
	if !strings.Contains(path, ":") {
		handler, exists := handlerTable[path]
		if !exists {
			statusCode = http.StatusNotFound
			responseBuff = []byte(fmt.Sprintf("not found handler. path: %s", path))
			return
		}
		buff, err := handler(r)
		if err != nil {
			return
		}
		statusCode = http.StatusOK
		responseBuff = buff
		return
	}

	paths := strings.Split(r.URL.Path, "/")

	responseBuff = []byte(strings.Join(paths, ", "))
	return statusCode, responseBuff
}

func rootHandler(r *http.Request) ([]byte, error) {
	return []byte("hello world!!!"), nil
}

func usersHandler(r *http.Request) ([]byte, error) {
	return []byte("users"), nil
}
