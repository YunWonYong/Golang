package main

import "fmt"

func main() {
	fmt.Println("=====interfaceTpyeAssertion1=====")
	interfaceTpyeAssertion1()
	fmt.Println("=====interfaceTpyeAssertion2=====")
	fmt.Println(interfaceTpyeAssertion2(1))
	fmt.Println(interfaceTpyeAssertion2('1'))
	fmt.Println(interfaceTpyeAssertion2(2.2))
	fmt.Println(interfaceTpyeAssertion2("Hello"))
	fmt.Println(interfaceTpyeAssertion2([]string{}))
	fmt.Println(interfaceTpyeAssertion2([...]string{"bye", "hi", "hello"}))
	fmt.Println(interfaceTpyeAssertion2(nil))
}

func interfaceTpyeAssertion1() {
	var x interface{} = 3
	a, isInt8 := x.(int8)
	b, isInt16 := x.(int16)
	c, isInt32 := x.(int32)
	d, isInt64 := x.(int64)
	e, isInt := x.(int)
	f, isUint8 := x.(uint8)
	g, isUint16 := x.(uint16)
	h, isUint32 := x.(uint32)
	i, isUint64 := x.(uint64)
	j, isUint := x.(uint)
	k, isFloat64 := x.(float64)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
	fmt.Println(isInt8, isInt16, isInt32, isInt64, isInt, isUint8, isUint16, isUint32, isUint64, isUint, isFloat64)
}

func interfaceTpyeAssertion2(i interface{}) string {
	switch x := i.(type) {
	case int:
		return fmt.Sprintf("case %T value: %#v, memory: %#v", x, i, &i)
	case string:
		return fmt.Sprintf("case %T value: %#v, memory: %#v", x, i, &i)
	case int32:
		return fmt.Sprintf("case %T value: %#v, memory: %#v", x, i, &i)
	case float64:
		return fmt.Sprintf("case %T value: %#v, memory: %#v", x, i, &i)
	case interface{}:
		return fmt.Sprintf("case %T value: %#v, memory: %#v", x, i, &i)
	default:
		return fmt.Sprintf("type!!!?? %s value: %#v, %#v", x, i, &i)
	}
}
