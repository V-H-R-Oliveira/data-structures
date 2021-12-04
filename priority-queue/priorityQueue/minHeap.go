package priorityqueue

import (
	"errors"
	"fmt"
	"math"
)

type MinBinHeap struct {
	Elements []*Node
}

func NewHeap() PriorityQueue {
	return &MinBinHeap{
		Elements: []*Node{},
	}
}

func (heap *MinBinHeap) bubbleUp(target *Node, targetIdx int) int {
	currentIdx, adjustedIdx := targetIdx, targetIdx

	for currentIdx > 0 {
		parentIdx := int(math.Round(float64(currentIdx)/2) - 1)
		parent := heap.Elements[parentIdx]

		if parent.Value > target.Value {
			heap.Elements[parentIdx], heap.Elements[currentIdx] = heap.Elements[currentIdx], heap.Elements[parentIdx]
			adjustedIdx = currentIdx
		}

		currentIdx = parentIdx
	}

	return adjustedIdx
}

func (heap *MinBinHeap) bubbleDown(targetIdx int) int {
	current, last := targetIdx, len(heap.Elements)-1

	for {
		leftIdx, rightIdx := 2*current+1, 2*current+2

		if leftIdx > last || rightIdx > last {
			break
		}

		left, right := heap.Elements[leftIdx], heap.Elements[rightIdx]

		if left.Value <= right.Value {
			heap.Elements[current], heap.Elements[leftIdx] = heap.Elements[leftIdx], heap.Elements[current]
			current = leftIdx
		} else {
			heap.Elements[current], heap.Elements[rightIdx] = heap.Elements[rightIdx], heap.Elements[current]
			current = rightIdx
		}
	}

	return current
}

func (heap *MinBinHeap) Insert(node *Node) error {
	if heap.IsEmpty() {
		heap.Elements = append(heap.Elements, node)
		return nil
	}

	heap.Elements = append(heap.Elements, node)
	targetIdx := len(heap.Elements) - 1
	heap.bubbleUp(node, targetIdx)

	return nil
}

func (heap *MinBinHeap) Poll() *Node {
	if heap.IsEmpty() {
		return nil
	}

	head := heap.Elements[0]
	lastIdx := len(heap.Elements) - 1

	heap.Elements[0], heap.Elements[lastIdx] = heap.Elements[lastIdx], heap.Elements[0]
	heap.Elements = heap.Elements[:lastIdx]

	lastIdx = len(heap.Elements) - 1
	targetIdx := heap.bubbleDown(0)

	leftChildIdx, rightChildIdx := 2*targetIdx+1, 2*targetIdx+2
	isOutBounds := leftChildIdx > lastIdx && rightChildIdx > lastIdx

	if !isOutBounds {
		needBubbleUpRight := false

		if rightChildIdx <= lastIdx {
			needBubbleUpRight = heap.Elements[targetIdx].Value > heap.Elements[rightChildIdx].Value
		}

		if heap.Elements[targetIdx].Value > heap.Elements[leftChildIdx].Value {
			heap.bubbleUp(heap.Elements[leftChildIdx], leftChildIdx)
		} else if needBubbleUpRight {
			heap.bubbleUp(heap.Elements[rightChildIdx], rightChildIdx)
		}
	}

	return head
}

func (heap *MinBinHeap) Peak() *Node {
	if heap.IsEmpty() {
		return nil
	}

	return heap.Elements[0]
}

func (heap *MinBinHeap) Remove(node *Node) error {
	if heap.IsEmpty() {
		return errors.New("the heap is empty")
	}

	nodeIdx := findByIndex(heap.Elements, node)

	if nodeIdx == -1 {
		return fmt.Errorf("element %d not found", node.Value)
	}

	lastIdx := len(heap.Elements) - 1

	heap.Elements[nodeIdx], heap.Elements[lastIdx] = heap.Elements[lastIdx], heap.Elements[nodeIdx]
	heap.Elements = heap.Elements[:lastIdx]

	lastIdx = len(heap.Elements) - 1

	targetIdx := heap.bubbleUp(heap.Elements[nodeIdx], nodeIdx)
	leftChildIdx, rightChildIdx := 2*targetIdx+1, 2*targetIdx+2
	isOutBounds := leftChildIdx > lastIdx && rightChildIdx > lastIdx

	if !isOutBounds {
		needBubbleDownRight := false

		if rightChildIdx <= lastIdx {
			needBubbleDownRight = heap.Elements[targetIdx].Value > heap.Elements[rightChildIdx].Value
		}

		needBubbleDownLeft := heap.Elements[targetIdx].Value > heap.Elements[leftChildIdx].Value

		if needBubbleDownLeft || needBubbleDownRight {
			heap.bubbleDown(targetIdx)
		}
	}

	return nil
}

func (heap *MinBinHeap) IsEmpty() bool {
	return len(heap.Elements) == 0
}

func (heap *MinBinHeap) Print() {
	if heap.IsEmpty() {
		return
	}

	for _, node := range heap.Elements {
		fmt.Printf("%d ", node.Value)
	}

	fmt.Println()
}

func (heap *MinBinHeap) Size() int {
	return len(heap.Elements)
}

func (heap *MinBinHeap) GetElements() []*Node {
	return heap.Elements
}