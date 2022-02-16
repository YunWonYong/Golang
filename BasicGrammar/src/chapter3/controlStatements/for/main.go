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
	fmt.Println("=====rangeFor=====")
	rangeFor()
	fmt.Println("=====rangeFor2=====")
	rangeFor2()
	fmt.Println("=====stringRangeFor=====")
	stringRangeFor()
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

func rangeFor() {
	arr := [3]string{"cat", "dog", "bird"}
	fmt.Printf("arr variable: %#v\n", arr)
	for index, element := range arr {
		fmt.Printf("arr index[%d] = %#v\n", index, element)
	}
}

func rangeFor2() {
	arr := [3]string{"cat", "dog", "bird"}
	fmt.Printf("arr variable: %#v\n", arr)
	index := 0
FirstFor:
	for _ /* 생략 가능 */, element := range arr {
		for {
			fmt.Printf("arr index[%d] = %#v\n", index, element)
			index++
			continue FirstFor
		}
	}
}

func stringRangeFor() {
	str := "가a나b다cㄱㄷㅏㅣ"

	for index, at := range str {
		fmt.Printf("at index[%02d], string[%s] type[%T] unicode[%#v]\n", index, string(at), at, at)
	}
}
