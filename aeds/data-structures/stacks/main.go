package main

import "fmt"

type stack struct {
	items []int
	size  int
}

func NewStack() *stack {
	return &stack{size: 0}
}

func (stack *stack) push(data int) {
	stack.items = append(stack.items, data)
	stack.size++
}

func (stack *stack) pop() int {
	stack.size--
	top := stack.items[stack.size]
	stack.items = stack.items[:stack.size]
	return top
}

func main() {
	myStack := NewStack()
	fmt.Println(myStack)
	myStack.push(3)
	myStack.push(2)
	myStack.push(4)
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
	fmt.Println(myStack)

}
