package main

import (
	"fmt"
	. "github.com/p4ali/go.util/stack"
)

func main() {
	stack := new(Stack)

	stack.Push(1)
	stack.Push("Hi")
	stack.Push(1.0)

	for stack.Len()>0 {
		fmt.Println(stack.Pop())
	}
}