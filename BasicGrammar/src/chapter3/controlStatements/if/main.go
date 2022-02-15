package main

import . "fmt"

func main() {

	Println("=====baseIfStatement=====")
	baseIfStatement()
	Println("=====baseIfElseStatement=====")
	baseIfElseStatement()
	Println("=====initializationIfStatement=====")
	initializationIfStatement()
	Println("=====ifStatementGotoApplyExample=====")
	ifStatementGotoApplyExample()
	Println("=====ifStatementDeferApplyExample=====")
	ifStatementDeferApplyExample()
}

func baseIfStatement() {
	i := 2
	if i%2 == 0 {
		Println("i는 짝수 입니다.")
		return
	}
	Println("i는 홀수 입니다.")
}

func baseIfElseStatement() {
	i := 21
	if i%5 == 0 {
		Printf("i(%d)는 5의 배수", i)
	} else if i%4 == 0 {
		Printf("i(%d)는 4의 배수", i)
	} else if i%2 == 0 {
		Printf("i(%d)는 2의 배수", i)
	} else {
		Printf("i(%d)는 3의 배수", i)
	}
	Println()
}

func initializationIfStatement() {
	if j := 2; j%3 == 0 {
		Println("if j:", j)
	} else if i := 3; i%2 == 0 {
		Printf("else if i: %d, j: %d\n", i, j)
	} else {
		Printf("else i: %d, j: %d\n", i, j)
	}

	//Println("if else out line j:", j) // error

	if i := 1; i%2 > 0 {
		Println("i는 홀수 입니다.")
		return
	}
	Println("i는 짝수 입니다.")
}

func ifStatementGotoApplyExample() {
	if i := 1; i%2 > 0 {
		Println("i는 홀수 입니다.")
		goto Even
	}
	Println("출력 해보거라~~")
Even: //label
	if i := 2; i%2 == 0 {
		Println("i는 짝수 입니다.")
		return
	}
}

func ifStatementDeferApplyExample() {
	if i := 1; i%2 > 0 {
		defer Println("i는 홀수 입니다.")
	}
	Println("뭐가 먼저 출력 될까?")
}
