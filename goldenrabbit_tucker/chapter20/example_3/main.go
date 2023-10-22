package main

import (
	"fmt"
        "math/rand"
        "time"
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


func (c *Cat) Eat(food string) {
	fmt.Printf("고양이가 %s을(를) 먹어요.\n", food)
}

func (d *Cat) Sleep() {
	fmt.Println("고양이가 골골 소리를 내며 자고 있어요.")
}

func (d *Cat) Bark() {
	fmt.Println("고양이가 야옹 소리로 짖고 있어요.")
}


const RAN = 10

func main() {
	var animals [RAN]Animal

        seed := time.Now().UnixNano()
        randSource := rand.NewSource(seed)
        rand := rand.New(randSource)
	for i := 0; i < RAN; i++ {
		if rand.Intn(RAN) % 2 == 0 {
			animals[i] = new(Cat)
			continue
		}
		animals[i] = new(Dog)
	}

	for i, animal := range animals {
		fmt.Printf("index: %d, animalType: %s\n", i, animalTypePrint(animal))
	}
}

func animalTypePrint(animal Animal) string {
	msg := "????"
	switch animal.(type) {
		case *Dog:
			msg = "Dog"
		case *Cat:
			msg = "Cat"
	}

	return msg
}

