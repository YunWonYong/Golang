package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 1번
		querys := r.URL.Query()
		// 2번
		name := querys.Get("name")
		ageStr := querys.Get("age")
		ageStr2 := querys.Get("Age")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("name: %s, age: %s, Age: %s", name, ageStr, ageStr2)))
	})

	http.ListenAndServe(":3000", nil)
}
