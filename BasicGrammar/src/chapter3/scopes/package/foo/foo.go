package foo

const (
	A   = 1
	abc = "abc"
)

var (
	m = 256
	N = 512
)

func DoSomething() {
	doSomething()
}

func doSomething() {
	println(abc, m)
}
