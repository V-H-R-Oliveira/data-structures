package priorityqueue

import "math"

func findByIndex(heap IHeap, target *Node) int {
	heapElements := heap.GetElements()

	if len(heapElements) == 0 {
		return -1
	}

	for idx, element := range heapElements {
		if element.Value == target.Value {
			return idx
		}
	}

	return -1
}

func bubbleDown(heap BinHeap, targetIdx int) int {
	heapElements := heap.GetElements()
	current, last := targetIdx, len(heapElements)-1

	for {
		leftIdx, rightIdx := 2*current+1, 2*current+2

		if leftIdx > last || rightIdx > last {
			break
		}

		left, right := heapElements[leftIdx], heapElements[rightIdx]

		if heap.ShouldBubbleDown(left, right) {
			heapElements[current], heapElements[leftIdx] = heapElements[leftIdx], heapElements[current]
			current = leftIdx
		} else {
			heapElements[current], heapElements[rightIdx] = heapElements[rightIdx], heapElements[current]
			current = rightIdx
		}
	}

	return current
}

func bubbleUp(heap BinHeap, target *Node, targetIdx int) int {
	heapElements := heap.GetElements()
	currentIdx, adjustedIdx := targetIdx, targetIdx

	for currentIdx > 0 {
		parentIdx := int(math.Round(float64(currentIdx)/2) - 1)
		parent := heapElements[parentIdx]

		if heap.ShouldBubbleUp(parent, target) {
			heapElements[parentIdx], heapElements[currentIdx] = heapElements[currentIdx], heapElements[parentIdx]
			adjustedIdx = currentIdx
		}

		currentIdx = parentIdx
	}

	return adjustedIdx
}
