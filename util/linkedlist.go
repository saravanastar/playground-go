package util

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (list *LinkedList[T]) Push(data T) {
	dataNode := Node[T]{Data: data}
	dataNode.Next = list.Tail
	dataNode.Prev = list.Tail.Prev
	list.Tail.Prev.Next = &dataNode
	list.Tail.Prev = &dataNode
	list.Size++
}

func (list *LinkedList[T]) PollLast() *T {
	if list.Size == 0 {
		return nil
	}
	dataNode := list.Tail.Prev
	dataNode.Prev.Next = dataNode.Next
	dataNode.Next.Prev = dataNode.Prev
	list.Size--
	return &dataNode.Data
}

func (list *LinkedList[T]) Poll() *T {
	if list.Size == 0 {
		return nil
	}
	dataNode := list.Head.Next
	list.Head.Next = dataNode.Next
	dataNode.Next.Prev = dataNode.Prev
	list.Size--
	return &dataNode.Data
}

func (list *LinkedList[T]) Length() int {
	return list.Size
}

func NewIntLinkedList() *LinkedList[int] {
	linkedList := LinkedList[int]{}
	headNode := Node[int]{}
	tailNode := Node[int]{}
	linkedList.Head = &headNode
	linkedList.Tail = &tailNode
	headNode.Next = &tailNode
	tailNode.Prev = &headNode
	return &linkedList
}
