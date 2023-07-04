package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdio = bufio.NewReader(os.Stdin)

func main() {
	randomN := GetRandomNumber(100)

	fmt.Println(randomN)
	tryCnt := 0
	msg := "1~100 랜덤 숫자 맞추기 게임"
	for {
		fmt.Println(msg)
		tryCnt++
		n, err := NumberInput()
		if err != nil {
			fmt.Println("숫자만 입력 가능합니다.")
			continue
		} else if n == randomN {
			fmt.Printf("축하합니다. 숫자를 맞추셨습니다. 시도횟수: %d\n", tryCnt)
			return
		}

		if n > randomN {
			msg = "입력하신 숫자가 더 큽니다."
			continue
		}

		msg = "입력하신 숫자가 더 작습니다."
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
