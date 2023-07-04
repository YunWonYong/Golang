package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const START_MONEY = 1000

var stdio = bufio.NewReader(os.Stdin)

func main() {
	msg := "1 ~ 5 숫자 입력하여 slot 돌려"
	money := START_MONEY
	for {
		if money <= 0 || money >= 5000 {
			fmt.Printf("ㅃ2 %d\n", money)
			return
		}
		fmt.Println(msg)
		slotNumber := GetRandomNumber(5)

		inputNumber, e := NumberInput()
		if e != nil {
			msg = "숫자만 입력"
			continue
		} else if inputNumber <= 0 || inputNumber > 5 {
			msg = "1 ~ 5 숫자만 입력"
			continue
		}

		if inputNumber == slotNumber {
			msg = "올..."
			money += 500
		} else {
			msg = "못 맞췄지롱."
			money -= 100
		}

		msg = fmt.Sprintf("%s 돈: %d원", msg, money)
	}
}

func GetRandomNumber(baseNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(baseNumber) + 1
}

func NumberInput() (n int, e error) {
	_, e = fmt.Scanln(&n)

	if e != nil {
		stdio.ReadString('\n')
	}

	return
}
