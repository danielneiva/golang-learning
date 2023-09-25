package main

import "fmt"

type node struct {
	next     *node
	previous *node
	data     int
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func NewLinkedList() linkedList {
	return linkedList{length: 0}
}

func (list *linkedList) prepend(node *node) {
	if list.length == 0 {
		list.head = node
		list.head.previous = list.head
		list.head.next = list.head
		list.tail = list.head
		list.tail.next = nil
		list.length++
		return
	}
	aux := list.head
	list.head = node
	list.head.next = aux
	aux.previous = list.head
	list.length++
}

func (list linkedList) printListData() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (list linkedList) printListDataReverse() {
	current := list.tail
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.previous
	}
	fmt.Println()
}

func main() {
	myList := NewLinkedList()
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	node5 := &node{data: 5}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.printListData()
	myList.printListDataReverse()
}
