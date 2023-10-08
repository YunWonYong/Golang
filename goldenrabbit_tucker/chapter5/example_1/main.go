package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v %v %v %v\n", nil, 'A', "A", 65)
	fmt.Printf("%T %T %T %T\n", nil, 'A', "A", 65)
	fmt.Printf("%t %t %t %t %t %t\n", nil, 'A', "A", 65, true, false)

}
