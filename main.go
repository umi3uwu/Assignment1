package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Println("before", list)
	handle(list)
	fmt.Println("after", append(list, 5))
}

func handle(slice []int) {
	newSlice := make([]int, len(slice))
	newSlice = append(slice, 5)
	fmt.Println("append", newSlice)
}
