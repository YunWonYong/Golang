package main

import "fmt"

func main() {
	var arr [5]int64 = [5]int64 { 1, 2, 3, 4 , 5 } // 명시적
	arr = [5]int64 { 1, 2, 3, 4 , 5 } // 묵시적
	arr = [...]int64 { 1, 2, 3, 4, 5 } // 배열의 길이는 모르겠고 원소 개수 만큼 할당 생성해줘~
	
	for index, element := range arr {
		fmt.Printf("index: %d, element: %d\n", index, element)
	}
}