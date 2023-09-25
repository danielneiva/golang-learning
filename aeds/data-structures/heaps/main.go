package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (heap *MaxHeap) Insert(key int) {
	heap.array = append(heap.array, key)
	heap.maxHeapifyUp(len(heap.array) - 1)
}

func (heap *MaxHeap) Extract() int {
	extracted := heap.array[0]
	lastIndex := len(heap.array) - 1

	if len(heap.array) == 0 {
		fmt.Println("Can't extract, array size is 0")
		return -1
	}
	heap.array[0] = heap.array[lastIndex]
	heap.array = heap.array[:lastIndex]

	heap.maxHeapifyDown(0)
	return extracted
}

func (heap *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(heap.array) - 1
	leftChild, rightChild := left(index), right(index)
	childToCompare := 0
	for leftChild <= lastIndex {
		if leftChild == lastIndex ||
			heap.array[leftChild] > heap.array[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}
		if heap.array[index] >= heap.array[childToCompare] {
			return
		}
		heap.swap(index, childToCompare)
		index = childToCompare
		leftChild, rightChild = left(index), right(index)
	}
}

func (heap *MaxHeap) maxHeapifyUp(index int) {
	for heap.array[parent(index)] < heap.array[index] {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (heap *MaxHeap) swap(i1, i2 int) {
	heap.array[i1], heap.array[i2] = heap.array[i2], heap.array[i1]
}

func main() {
	max := &MaxHeap{}
	fmt.Println(max)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		max.Insert(v)
		fmt.Println(max)
	}

	for i := 0; i < 5; i++ {
		max.Extract()
		fmt.Println(max)
	}

}
