package linkedlist

func New[T nodeValueType](initValues ...T) *linkedList[T] {
	ll := new(linkedList[T])

	for _, value := range initValues {
		ll.Push(value)
	}
	return ll
}
