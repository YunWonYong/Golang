package main

import "fmt"

func main() {
	fmt.Println("=====baseForStatement=====")
	baseForStatement()
	fmt.Println("=====infinityLoopAndBreak=====")
	infinityLoopAndBreak()
	fmt.Println("=====initStatementFor=====")
	initStatementFor()
	fmt.Println("=====conditionExpressionFor=====")
	conditionExpressionFor()
	fmt.Println("=====postStatementAndContinueFor=====")
	postStatementAndContinueFor()
	fmt.Println("=====postStatementAndContinueFor2=====")
	postStatementAndContinueFor2()
}

func baseForStatement() {
	defer fmt.Println()
	for i := 1; i < 11; i++ {
		fmt.Printf("%d ", i)
	}
}

func infinityLoopAndBreak() {
	breakPoint := 0
	for {
		if breakPoint == 100 {
			fmt.Println("break!!!")
			break
		}
		breakPoint++
	}
}

func initStatementFor() {
	for i := 0; ; {
		if i == 2 {
			fmt.Printf("i(%d) stop!\n", i)
			break
		}
		i++
	}
}

func conditionExpressionFor() {
	i := 2
	defer fmt.Println("defer i: ", i) // ???? 왜 2??
	for i < 100 {
		i++
	}
	fmt.Println("loop end i: ", i)
}

func postStatementAndContinueFor() {
	defer fmt.Println()
	i := 0
	for ; i < 101; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

func postStatementAndContinueFor2() {
	defer fmt.Println()
	i := 100
	for ; i > 0; i -= 2 {
		if i%2 != 0 {
			fmt.Print("continue!!! ")
			continue
		}
		fmt.Printf("%d ", i)
	}
}
