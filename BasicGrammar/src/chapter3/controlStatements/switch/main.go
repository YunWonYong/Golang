package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====baseSwitchStatement=====")
	baseSwitchStatement()
	fmt.Println("=====switchDefaultCase=====")
	switchDefaultCase()
	fmt.Println("=====switchCaseTypeCheckAndFallthrough=====")
	switchCaseTypeCheckAndFallthrough()
	fmt.Println("=====caseMultiValueFourSeasonsPrint=====")
	caseMultiValueFourSeasonsPrint()
	fmt.Println("=====switchExpressionCase=====")
	switchExpressionCase(5)
	switchExpressionCase(20)
	switchExpressionCase(31)
}

func baseSwitchStatement() {
	/**
	 go의 switch문은 기본으로 case 한 개당 break가 발생한다.
	 break 예약어가 없다해서 계속 진행하지 않고 case 순서 상관 없음
	**/
	n := 10
	switch n {
	case 1:
		fmt.Println("n = 1")
		// break => redundant
	case 2:
		fmt.Println("n = 2")
	case 3:
		fmt.Println("n = 3")
	case 4:
		fmt.Println("n = 4")
	case 10:
		fmt.Println("n = 10")
	case 5:
		fmt.Println("n = 5")
	case 6:
		fmt.Println("n = 6")
	case 7:
		fmt.Println("n = 7")
	case 8:
		fmt.Println("n = 8")
	case 9:
		fmt.Println("n = 9")
	default:
		fmt.Println("n = NaN")
	}
}

func switchDefaultCase() {
	switch {
	case true:
		fmt.Println("1. true")
	default:
		fmt.Println("1. false")
	}

	switch {
	case false:
		fmt.Println("2. case false")
	default:
		fmt.Println("2. default false")
	}

	switch {
	case false:
		fmt.Println("3. case false")
	case true:
		fmt.Println("3. case true")
	default:
		fmt.Println("3. default false")
	}

	x, y := 1, 1

	switch {
	case false:
		fmt.Println("4. case false")
	case x == y:
		fmt.Println("4. case x == y")
	case true:
		fmt.Println("4. case true")
	default:
		fmt.Println("4. default false")
	}
}

func switchCaseTypeCheckAndFallthrough() {
	arr := [13]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 13}
	msg := ""
	flag := true
	for _, element := range arr {
		flag = true
		switch element {
		//case "1": // error
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			msg = "봄"
			goto strAppend
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			msg = "여름"
			goto strAppend
		case 9:
			fallthrough
		case 10:
			fallthrough
		case 11:
			msg = "가을"
			goto strAppend
		case 12.0: // 실수에 소수가 없을 시 정수로 변환하여 맵핑해줌
			fallthrough
		case 1:
			fallthrough
		case 2:
			msg = "겨울"
			goto strAppend
		default:
			msg = "없음!!"
			flag = false
		}
	strAppend:
		if flag {
			msg += "입니다."
		}
		fmt.Printf("%02d월은 %s\n", element, msg)
	}

}

func caseMultiValueFourSeasonsPrint() {
	arr := [13]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 13}
	for _, element := range arr {
		switch element {
		case 1, 2.0, 12.0:
			fmt.Printf("%02d월은 겨울입니다.\n", element)
		case 3, 4, 5.0:
			fmt.Printf("%02d월은 봄입니다.\n", element)
		case 6, 7, 8.0:
			fmt.Printf("%02d월은 여름입니다.\n", element)
		case 9, 10.00000, 11:
			fmt.Printf("%02d월은 가을입니다.\n", element)
		default:
			fmt.Printf("%02d월은 없음!!\n", element)
		}
	}
}

func switchExpressionCase(n int) {
	switch {
	case n > 0 && n < 11:
		fmt.Println("n은 1 ~ 10 사이다.")
	case n > 10 && n < 21:
		fmt.Println("n은 11 ~ 20 사이다.")
	case n > 20 && n < 31:
		fmt.Println("n은 21 ~ 30 사이다.")
	default:
		fmt.Println("n은 0보다 작거나 30보다 크다.")
	}
}
