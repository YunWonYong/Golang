package main

import (
	"fmt"
	_ "math"
)

func main() {
	c1 := 1.0 + 3i
	fmt.Printf("c1 value: %#v, type: %T, real: %g, imag: %g\n", c1, c1, real(c1), imag(c1))
	c2 := complex(1.0, 3)
	fmt.Printf("c2 value: %#v, type: %T, real: %g, imag: %g\n", c2, c2, real(c2), imag(c2))
	var c3 complex64 = 1.0 + 3i
	fmt.Printf("c3 value: %#v, type: %T, real: %g, imag: %g\n", c3, c3, real(c3), imag(c3))
	c4 := complex64(1.0 + 3i)
	fmt.Printf("c4 value: %#v, type: %T, real: %g, imag: %g\n", c4, c4, real(c4), imag(c4))
}
