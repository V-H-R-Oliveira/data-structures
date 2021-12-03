package main

import (
	"fmt"
	priorityqueue "priorityQueue/m/priorityQueue"
)

func main() {
	values := []int{1, 5, 1, 8, 6, 2, 2, 13, 12, 11, 7, 2, 15, 3, 10}
	heap := priorityqueue.NewHeap()

	for _, value := range values {
		heap.Insert(priorityqueue.NewNode(value))
	}

	fmt.Println("Poll:", heap.Poll())
	fmt.Println("Remove:", heap.Remove(priorityqueue.NewNode(12)))
	fmt.Println("Remove:", heap.Remove(priorityqueue.NewNode(3)))
	fmt.Println("Poll:", heap.Poll())
	fmt.Println("Remove:", heap.Remove(priorityqueue.NewNode(6)))
	heap.Print()
}
