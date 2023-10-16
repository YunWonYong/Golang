package linkedList

import "fmt"

type linkedList[T nodeValueType] struct {
	head *node[T]
	tail *node[T]
}

func (ll *linkedList[T]) PushAll(values ...T) {
	for _, value := range values {
		ll.Push(value)
	}
}

func (ll *linkedList[T]) Push(value T) {
	newNode := new(node[T]).setValue(value)
	pushFn := ll.append
	if ll.IsEmpty() {
		pushFn = ll.setHead
	} else if ll.head != nil && ll.tail == nil {
		pushFn = ll.setTail
	}

	pushFn(newNode)
}

func (ll *linkedList[T]) append(newNode *node[T]) {
	prev := ll.tail
	ll.tail = newNode.setPrev(prev)

	prev.setNext(ll.tail)

}

func (ll *linkedList[T]) setHead(newNode *node[T]) {
	if ll.tail != nil && ll.tail == newNode {
		ll.head = ll.tail.setPrev(nil)
		ll.tail = nil
		return
	}

	n := newNode
	if newNode != nil {
		n = newNode.setPrev(nil)
	}
	ll.head = n
}

func (ll *linkedList[T]) setTail(newNode *node[T]) {
	if newNode != nil && ll.tail == nil {
		ll.head.setNext(newNode)
		newNode.setPrev(ll.head)
	} else if ll.head == newNode {
		ll.head.setNext(nil)
		newNode = nil
	} else if newNode != nil {
		newNode.setNext(nil)
	}

	ll.tail = newNode
}

func (ll *linkedList[T]) Pop() (t T, err error) {
	if ll.IsEmpty() {
		err = fmt.Errorf("linked list empty")
		return
	} else if ll.tail != nil {
		t = ll.tail.value
		ll.setTail(ll.tail.prev)
		return
	}
	t = ll.head.value
	ll.setHead(nil)
	return
}

func (ll *linkedList[T]) Unshift() (t T, err error) {
	if ll.head == nil {
		err = fmt.Errorf("linked list empty")
		return
	}

	t = ll.head.value
	ll.setHead(ll.head.next)

	return
}

func (ll *linkedList[T]) IsEmpty() bool {
	return ll.head == nil && ll.tail == nil
}
