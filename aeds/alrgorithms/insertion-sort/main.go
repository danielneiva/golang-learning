package main

import (
	"fmt"
	"io"
	"os"
)

func loadNumbers(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", filePath, err))
	}
	var line int
	for {
		_, err := fmt.Fscanf(fd, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			panic(fmt.Sprintf("Sacn failed %s: %v", filePath, err))
		}
		numbers = append(numbers, line)
	}
}

func printNumbers(numbers []int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
	fmt.Println()
}

func insertionSort(numbers []int) {
	length := len(numbers)
	for i := 0; i < length-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			for j := i; j > 0; j-- {
				if numbers[j] < numbers[j-1] {
					numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
				}
			}
		}
	}
}

func verifyIfIsSorted(numbers []int) bool {
	for i := range numbers[:len(numbers)-1] {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Loading numbers")
	numbers := loadNumbers("numbers.txt")
	fmt.Println("Sorting numbers")
	insertionSort(numbers)
	printNumbers(numbers)
	fmt.Println(verifyIfIsSorted(numbers))
}
