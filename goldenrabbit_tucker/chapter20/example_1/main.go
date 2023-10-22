package main

import (
	"fmt"
	"encoding/json"
)

type (
	A struct {
		Name string `json:"name"`
		Age int64 `json:"age"`
	}

	B struct {
		name string `json:"name"`
		age int64 `json:"age"`
	}
)
func main() {

	data := []byte("{\"name\": \"수빈\", \"age\": 100}")
	a := new(A)

	fmt.Printf("a Unmarshalling before\n%#+v\n", a)
	Unmarshalling(data, a)
	fmt.Printf("a Unmarshalling after\n%#+v\n", a)
	data = Marshalling(a)
	fmt.Printf("a Marshalling result: %s\n", string(data))

	b := new(B)
	fmt.Printf("a Marshalling data b Unmarshalling before\n%#+v\n", b)
        Unmarshalling(data, b)
	fmt.Printf("a Marshalling data b Unmarshalling after\n%#+v\n", b)
}

func Marshalling(obj interface{}) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return data
}

func Unmarshalling(data []byte, obj interface{}) {
	if err := json.Unmarshal(data, obj); err != nil {
		panic(err)
	}
}
