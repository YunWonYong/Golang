package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1번
	http.HandleFunc("/*/*", func(w http.ResponseWriter, r *http.Request) {
		paths := r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pattern: /*/*, paths: %s", paths)))
	})

	// 2번
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		paths := r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pattern: /*, paths: %s", paths)))
	})

	// 3번
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths := r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pattern: /, paths: %s", paths)))
	})

	// 4번
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		paths := r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pattern: /test, paths: %s", paths)))
	})

	http.ListenAndServe(":3000", nil)
}
