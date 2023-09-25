package main

import "fmt"

type node struct {
	data int
	next *node
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
	aux := list.head
	list.head = node
	list.head.next = aux
	if list.length == 0 {
		list.tail = aux
	}
	list.length++
}

func (list *linkedList) append(node *node) {
	if list.length == 0 {
		list.head = node
		list.tail = list.head
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.length++
}

func (list linkedList) hasValue(value int) bool {
	current := list.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (list linkedList) printListData() {
	current := list.head
	for list.length != 0 {
		fmt.Printf("%d ", current.data)
		current = current.next
		list.length--
	}
	fmt.Println()
}

func (list *linkedList) deleteByValue(value int) {
	if list.length == 0 {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return
	}
	current := list.head
	for current.next.data != value {
		current = current.next
		if current.next == nil {
			return
		}
	}
	current.next = current.next.next
	list.length--
}

func main() {
	myList := NewLinkedList()
	node1 := &node{data: 2}
	node2 := &node{data: 3}
	node3 := &node{data: 5}
	node4 := &node{data: 9}
	node5 := &node{data: 21}
	node6 := &node{data: 12}
	node7 := &node{data: 7}
	myList.append((node1))
	myList.append((node2))
	myList.append((node3))
	myList.append((node4))
	myList.append((node5))
	myList.append((node6))
	myList.prepend((node7))
	myList.printListData()
}
