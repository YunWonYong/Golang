package main

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func (node *Node) init(prev *Node, value interface{}, next *Node) *Node {
	node.prev = prev
	node.value = value
	node.next = next
	return node
}

func main() {
	var node1_1 *Node = new(Node) // 명시적 1번 방법
	node1_1.prev = nil
	node1_1.value = "Hello"

	var node1_2 *Node = &Node{ // 명시적 2번 방법
		prev:  node1_1,
		value: " World",
	}

	node2_1 := new(Node) // 묵시적 1번 방법
	node2_2 := &Node{}   // 묵시적 2번 방법

	node1_1.next = node1_2
	node1_2.next = node2_1

	node2_1.init(node1_2, 1992, node2_2)
	node2_2.init(node2_2, 1107, nil)

	fmt.Printf("%p %p %p %p\n", node1_1, node1_2, node2_1, node2_2)
}
