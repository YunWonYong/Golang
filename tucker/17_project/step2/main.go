package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("1~100 사이의 숫자를 입력해 주세요.")
	n, err := Input()

	if err != nil {
		fmt.Println("숫자만 입력 가능합니다.")
		return
	}

	fmt.Printf("입력하신 숫자는 %d 입니다.\n", n)
}

func Input() (n int, e error) {

	n, e = fmt.Scanln(&n)

	if e != nil {
		stdin.ReadString('\n')
	}

	return n, e
}
