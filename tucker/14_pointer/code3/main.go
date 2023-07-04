package main

import "fmt"

type Data struct {
	a int64
	b int
}

func main() {
	pointerInstanceCopy()
	valueInstanceCopy()
}

func pointerInstanceCopy() {
	a := new(Data)
	b := a
	c := a

	fmt.Printf("new function instance\n a: %p, b: %p, c: %p\n", a, b, c)

	d := &Data{}
	e := d
	f := d

	fmt.Printf("literal pointer instance\n d: %p, e: %p, f: %p\n", d, e, f)
}

func valueInstanceCopy() {
	a := Data{}
	b := a
	c := a

	fmt.Printf("literal instance\n a: %p, b: %p, c: %p\n", &a, &b, &c)
}
