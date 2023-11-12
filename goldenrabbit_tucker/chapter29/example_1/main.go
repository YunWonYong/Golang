package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		querys := r.URL.Query()
		name := querys.Get("name")
		ageStr := querys.Get("age")
		ageStr2 := querys.Get("Age")
		fmt.Println(name, ageStr, ageStr2)

		// response 1번
		w.WriteHeader(http.StatusOK)
		// response 2번
		w.Write([]byte(fmt.Sprintf("name: %s, age: %s, Age: %s", name, ageStr, ageStr2)))
	})

	http.ListenAndServe(":3000", nil)
}
