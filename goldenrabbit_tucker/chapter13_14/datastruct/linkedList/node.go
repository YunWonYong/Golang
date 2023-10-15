package linkedlist

type (
	nodeValueType interface {
		string | int64 | any
	}

	node[T nodeValueType] struct {
		prev  *node[T]
		value T
		next  *node[T]
	}
)

func (n *node[T]) setPrev(prev *node[T]) *node[T] {
	n.prev = prev
	return n
}

func (n *node[T]) setValue(value T) *node[T] {
	n.value = value
	return n
}

func (n *node[T]) setNext(next *node[T]) *node[T] {
	n.next = next
	return n
}
