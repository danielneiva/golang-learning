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
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Sacn failed %s: %v", filePath, err))
		}
		numbers = append(numbers, line)
	}
}

func main() {
	numbers := loadNumbers("numbers.txt")
	carry := 0
	for _, v := range numbers {
		carry += v
	}
	fmt.Println(carry)
}
