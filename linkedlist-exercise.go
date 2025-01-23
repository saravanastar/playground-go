package main

import (
	"fmt"

	"github.com/saravanastar/playground-go/util"
)

func main() {

	linkedList := util.NewIntLinkedList()
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Push(3)

	fmt.Println(*linkedList.Poll())
	fmt.Println(*linkedList.Poll())
	fmt.Println(*linkedList.Poll())
	// fmt.Println(*linkedList.Poll())
	linkedList.Push(4)
	linkedList.Push(5)
	fmt.Println(*linkedList.Poll())
	fmt.Println(*linkedList.Poll())
}
