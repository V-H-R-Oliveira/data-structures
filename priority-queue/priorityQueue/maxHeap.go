package priorityqueue

import (
	"errors"
	"fmt"
)

type MaxBinHeap struct {
	Elements []*Node
}

func NewMaxHeap() IHeap {
	return &MaxBinHeap{
		Elements: []*Node{},
	}
}

func (heap *MaxBinHeap) ShouldBubbleUp(source, target *Node) bool {
	return target.Value > source.Value
}

func (heap *MaxBinHeap) ShouldBubbleDown(left, right *Node) bool {
	return left.Value >= right.Value
}

func (heap *MaxBinHeap) Insert(node *Node) error {
	if heap.IsEmpty() {
		heap.Elements = append(heap.Elements, node)
		return nil
	}

	heap.Elements = append(heap.Elements, node)
	targetIdx := len(heap.Elements) - 1
	bubbleUp(heap, node, targetIdx)

	return nil
}

func (heap *MaxBinHeap) Poll() *Node {
	if heap.IsEmpty() {
		return nil
	}

	head := heap.Elements[0]
	lastIdx := len(heap.Elements) - 1

	heap.Elements[0], heap.Elements[lastIdx] = heap.Elements[lastIdx], heap.Elements[0]
	heap.Elements = heap.Elements[:lastIdx]

	lastIdx = len(heap.Elements) - 1
	targetIdx := bubbleDown(heap, 0)

	leftChildIdx, rightChildIdx := 2*targetIdx+1, 2*targetIdx+2
	isOutBounds := leftChildIdx > lastIdx && rightChildIdx > lastIdx

	if !isOutBounds {
		needBubbleUpRight := false

		if rightChildIdx <= lastIdx {
			needBubbleUpRight = heap.Elements[targetIdx].Value < heap.Elements[rightChildIdx].Value
		}

		if heap.Elements[targetIdx].Value < heap.Elements[leftChildIdx].Value {
			bubbleUp(heap, heap.Elements[leftChildIdx], leftChildIdx)
		} else if needBubbleUpRight {
			bubbleUp(heap, heap.Elements[rightChildIdx], rightChildIdx)
		}
	}

	return head
}

func (heap *MaxBinHeap) Remove(node *Node) error {
	if heap.IsEmpty() {
		return errors.New("the heap is empty")
	}

	nodeIdx := findByIndex(heap, node)

	if nodeIdx == -1 {
		return fmt.Errorf("element %d not found", node.Value)
	}

	lastIdx := len(heap.Elements) - 1

	heap.Elements[nodeIdx], heap.Elements[lastIdx] = heap.Elements[lastIdx], heap.Elements[nodeIdx]
	heap.Elements = heap.Elements[:lastIdx]

	lastIdx = len(heap.Elements) - 1

	targetIdx := bubbleUp(heap, heap.Elements[nodeIdx], nodeIdx)
	leftChildIdx, rightChildIdx := 2*targetIdx+1, 2*targetIdx+2
	isOutBounds := leftChildIdx > lastIdx && rightChildIdx > lastIdx

	if !isOutBounds {
		needBubbleDownRight := false

		if rightChildIdx <= lastIdx {
			needBubbleDownRight = heap.Elements[targetIdx].Value < heap.Elements[rightChildIdx].Value
		}

		needBubbleDownLeft := heap.Elements[targetIdx].Value < heap.Elements[leftChildIdx].Value

		if needBubbleDownLeft || needBubbleDownRight {
			bubbleDown(heap, targetIdx)
		}
	}

	return nil
}

func (heap *MaxBinHeap) GetElements() []*Node {
	return heap.Elements
}

func (heap *MaxBinHeap) Size() int {
	return len(heap.Elements)
}

func (heap *MaxBinHeap) IsEmpty() bool {
	return len(heap.Elements) == 0
}

func (heap *MaxBinHeap) Peak() *Node {
	if heap.IsEmpty() {
		return nil
	}

	return heap.Elements[0]
}

func (heap *MaxBinHeap) Print() {
	if heap.IsEmpty() {
		return
	}

	for _, node := range heap.Elements {
		fmt.Printf("%d ", node.Value)
	}

	fmt.Println()
}
