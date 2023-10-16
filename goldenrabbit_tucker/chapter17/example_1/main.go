package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	RANDOM_NUMBER_GENERATE_BOUNDARY int = 100

	INPUT_NUMBER_MESSAGE         string = "숫자값을 입력하세요:"
	INPUT_NUMBER_ERROR_MESSAGE   string = "숫자만 입력하세요."
	INPUT_NUMBER_LESS_MESSAGE    string = "입력하신 숫자가 더 작습니다."
	INPUT_NUMBER_GREATER_MESSAGE string = "입력하신 숫자가 더 큽니다."
	CLEAR_MESSAGE                string = "축하합니다. 숫자를 맞추셨습니다. 시도횟수: %d번"
)

func main() {
	targetNum := getRandomNumber(RANDOM_NUMBER_GENERATE_BOUNDARY)
	inputCnt := 0
	msgNum := 1
	for msgNum != 3 {
		inputCnt++
		inputNum, err := inputNumber()
		if err != nil {
			continue
		}

		print(getNumberCheckState(targetNum, inputNum), inputCnt)
	}
}

func getRandomNumber(boundary int) int {
	seed := time.Now().UnixNano()
	randSource := rand.NewSource(seed)
	rand := rand.New(randSource)
	return rand.Intn(boundary)
}

func inputNumber() (int, error) {
	fmt.Print(INPUT_NUMBER_MESSAGE)
	inputNum := 0
	_, err := fmt.Scanf("%d\n", &inputNum)

	if err != nil {
		fmt.Println(INPUT_NUMBER_ERROR_MESSAGE)
		inputStreamFlush()
		return 0, err
	}
	return inputNum, nil
}

func inputStreamFlush() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func getNumberCheckState(targetNumber, inputNumber int) int {
	if targetNumber < inputNumber {
		return 1
	} else if targetNumber > inputNumber {
		return 2
	}
	return 3
}

func print(messageNumber, inputNum int) {
	msg := ""
	switch messageNumber {
	case 1:
		msg = INPUT_NUMBER_LESS_MESSAGE
	case 2:
		msg = INPUT_NUMBER_GREATER_MESSAGE
	case 3:
		msg = fmt.Sprintf(CLEAR_MESSAGE, inputNum)
	}

	if len(msg) == 0 {
		panic(fmt.Sprintf("not supported messageNumber: %d", messageNumber))
	}

	fmt.Println(msg)
}
