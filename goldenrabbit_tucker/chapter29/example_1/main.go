package main

import (
	// 1번
	"net/http"
)

func main() {

	// 2번
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// response 1번
		w.WriteHeader(http.StatusOK)
		// response 2번
		w.Write([]byte("hello world!!!"))
	})

	// 3번
	http.ListenAndServe(":3000", nil)
}
