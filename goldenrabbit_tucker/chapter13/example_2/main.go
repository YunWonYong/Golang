package main

import "fmt"

type Node struct {
	value interface{}
}

func main() {
	var node1 Node
	fmt.Println(node1.value)
	var node2 Node = Node{}
	node3 := Node{}

	fmt.Printf("%#+v %#+v %#+v\n", node1, node2, node3)
}
