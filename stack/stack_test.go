package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	data := stack.Pop()
	fmt.Println(*data)
	stack.Print()

	stack.Push(-1)
	stack.Print()

	data = stack.Pop()
	fmt.Println(*data)
	stack.Print()
	data = stack.Pop()
	fmt.Println(*data)
	stack.Print()
	data = stack.Pop()
	fmt.Println(data)
	stack.Print()
}
