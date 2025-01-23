package util

type Queue[T any] struct {
	elements []*T
	size     int
}

func (queue *Queue[T]) Enqueue(element T) bool {
	queue.elements = append(queue.elements, &element)
	queue.size++
	return true
}

func (queue *Queue[T]) Dequeue() *T {
	removedElement := queue.elements[0]
	queue.elements = queue.elements[1:]
	return removedElement
}

func (queue *Queue[T]) Peek() *T {
	if len(queue.elements) == 0 {
		return nil
	}

	return queue.elements[0]
}

func NewIntQueue() *Queue[int] {
	return &Queue[int]{elements: make([]*int, 0)}
}
