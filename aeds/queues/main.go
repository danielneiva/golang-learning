package main

import "fmt"

type queue struct {
	items []int
	size  int
}

func NewQueue() *queue {
	return &queue{size: 0}
}

func (queue *queue) enqueue(value int) {
	queue.items = append(queue.items, value)
	queue.size++
}

func (queue *queue) dequeue() int {
	queue.size--
	first := queue.items[0]
	queue.items = queue.items[1:]
	return first
}

func main() {
	myQueue := NewQueue()

	myQueue.enqueue(10)
	myQueue.enqueue(12)
	myQueue.enqueue(4)
	myQueue.enqueue(9)
	myQueue.enqueue(120)
	fmt.Println(myQueue)
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue)

}
