package main

import (
	"fmt"
)

func main() {
	//Go는 Enum이 없다.
	const (
		A = iota
		B
		C
		D
	)
	fmt.Printf("A = %d, B = %d, C = %d, D = %d\n", A, B, C, D)

	const (
		E = iota
		F = 1000
		G = "가나다라"
		H = iota // 1?
	)
	fmt.Printf("E = %d, F = %d, G = %s, H = %d\n", E, F, G, H)

	const I = iota
	const J = iota

	fmt.Printf("I == J ? %t\n", I == J)

	const (
		K = 1 + iota
		L = 1 + iota
		M = 1 + iota
		N = iota
	)

	fmt.Printf("K = %d, L = %d, M = %d, N = %d\n", K, L, M, N)
	/*
		I 상수의 iota와 J 상수의 iota는 관계가 없고
		const 블럭에서 여러개의 상수 선언 시 iota가 의미있다.
		iota는 쉽게 생각하면 const 블럭에서 선언한 변수의 index라 생각해도 괜찮다.
		iota에 연산한다 해서 iota의 값이 반영되지 않는다.
	*/
}
