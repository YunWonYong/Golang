package main

import (
	. "fmt"
	"sync"
)

var (
	slice = []int{}
	flag  = false
)

func main() {
	a := []int{1, 2, 3, 5}
	Printf("append after a value: %#v, %p, len: %d, cap: %d\n", a, &a, len(a), cap(a))
	a = append(a, 2)
	Printf("append before a value: %#v, %p, len: %d, cap: %d\n", a, &a, len(a), cap(a))
	var wg sync.WaitGroup
	go routine("1번 루틴", &wg)
	go routine("2번 루틴", &wg)
	go routine("3번 루틴", &wg)
	go routine("4번 루틴", &wg)

	for {
		if flag {
			wg.Wait()
			break
		}
	}
}

func routine(name string, wg *sync.WaitGroup) {
	wg.Add(1)
	wg.Done()
	defer func() {
		if str := recover(); str != nil {
			Println(str, "마무의리!!!", cap(slice), len(slice))
			flag = true
		}
	}()
	capNum := 0
	temp := 0
	for {
		temp = cap(slice)
		if capNum < temp {
			Printf("old cap : %d, extend cap: %d\n", capNum, temp)
			capNum = temp
		}
		slice = append(slice, 0)
		if cap(slice) > 100 {
			panic(name)
		}
	}
}
