package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName()) //go run main.go app.go 해야 실행됨
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
