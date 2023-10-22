package main

import (
	"fmt"
)

type (
	Animal interface {
		Eat(food string)
		Sleep()
		Bark()

	}
	Dog struct {
		Animal
	}

	Cat struct {
		Animal
	}
)

func (d *Dog) Eat(food string) {
	fmt.Printf("개가 %s을(를) 먹어요.\n", food)
}

func (d *Dog) Sleep() {
	fmt.Println("개가 쿨쿨 자고 있어요.")
}

func (d *Dog) Bark() {
	fmt.Println("개가 왈왈 소리로 짖고 있어요.")
}

func (c *Cat) Eat(food string) {
	fmt.Printf("고양이가 %s을(를) 먹어요.\n", food)
}

func (d *Cat) Sleep() {
	fmt.Println("고양이가 골골 소리를 내며 자고 있어요.")
}

func (d *Cat) Bark() {
	fmt.Println("고양이가 야옹 소리로 짖고 있어요.")
}

func main() {
	var animals [2]Animal

	animals[0] = new(Dog)
	animals[1] = new(Cat)

	for _, animal := range animals {
		animal.Eat("사료")
		animal.Bark()
		animal.Sleep()
	}
}

