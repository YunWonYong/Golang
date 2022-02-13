package main

import "fmt"

func main() {
	var sum func(a, b int, f func(x int)) = func(a, b int, f func(x int)) { f(a + b) }
	sum(5, 10, func(x int) { fmt.Println(x) })

	minus := func(a, b int, f func(x int)) (result int) {
		result = a - b
		f(result)
		return
	}

	result := minus(10, 6, func(x int) { fmt.Println("callbackFn: ", x) })
	fmt.Println("result: ", result)

	gcd := func(n, m int) int {
		for z := 0; ; {
			if m == 0 {
				return n
			}
			z = n
			n = m
			m = z % n
		}
	}
	lcd := func(n, m int, gcd func(n, m int) int) (int, int) {
		g := gcd(n, m)
		return g, n * (m / g)
	}

	func(a, b int, lcd func(a, b int, gcd func(a, b int) int) (int, int)) {
		fmt.Println(lcd(a, b, gcd))
	}(90, 1000000000000000000, lcd)
}
