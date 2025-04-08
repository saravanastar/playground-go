package sum

import (
	"container/list"
	"errors"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue[T any] struct {
	data *list.List
}

func NewQueue() *Queue[*TreeNode] {
	return &Queue[*TreeNode]{
		data: list.New(),
	}
}

func (queue *Queue[T]) Front() (T, error) {
	if queue.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	return queue.data.Front().Value.(T), nil
}

func (queue *Queue[T]) Push(val T) {
	queue.data.PushBack(val)
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.data.Len() == 0
}

func (queue *Queue[T]) Pop() (T, error) {
	if queue.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	elementToPop := queue.data.Front()
	queue.data.Remove(queue.data.Front())
	return elementToPop.Value.(T), nil
}

func sumOfNodes(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}

	queue := NewQueue()
	queue.Push(root)
	for !queue.IsEmpty() {
		currentNode, _ := queue.Pop()
		result += currentNode.Val
		if currentNode.Left != nil {
			queue.Push(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.Push(currentNode.Right)
		}
	}

	return result

}
