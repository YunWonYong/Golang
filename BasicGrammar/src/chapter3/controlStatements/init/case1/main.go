package main

import "fmt"

var (
	str string
)

func init() {
	str += "1"
}

func init() {
	str += "2"
}

func main() {
	fmt.Println(str)
}

func init() {
	str += "3"
}

func init() {
	str += "4"
}
