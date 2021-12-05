package priorityqueue_test

import (
	priorityqueue "priorityQueue/m/priorityQueue"
	"testing"
)

func TestMaxBinHeap(t *testing.T) {
	t.Run("Test isEmpty", func(t *testing.T) {
		t.Run("Given an empty heap, it should return true", func(t *testing.T) {
			heap := priorityqueue.NewMaxHeap()

			if !heap.IsEmpty() {
				t.Error("Expected an empty heap, got a filled heap")
			}
		})

		t.Run("Given a fulfilled heap, it should return false", func(t *testing.T) {
			values := []int{1, 5, 1, 8, 6, 2, 2, 13, 12, 11, 7, 2, 15, 3, 10}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			if heap.IsEmpty() {
				t.Error("Expected a filled heap, got an empty heap")
			}
		})
	})

	t.Run("Test Insert", func(t *testing.T) {
		t.Run("Given an input, the new heap should have the same length as the input after", func(t *testing.T) {
			values := []int{1, 5, 1, 8, 6, 2, 2, 13, 12, 11, 7, 2, 15, 3, 10}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			heapSize := heap.Size()
			inputSize := len(values)

			if heapSize != inputSize {
				t.Errorf("Expected %d, got %d\n", inputSize, heapSize)
			}
		})
	})

	t.Run("Test heap invariance property", func(t *testing.T) {
		t.Run("Test #1", func(t *testing.T) {
			values := []int{1, 5, 1, 8, 6, 2, 2, 13, 12, 11, 7, 2, 15, 3, 10}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			heapValues := heap.GetElements()
			expectedHeap := []int{15, 12, 13, 8, 11, 2, 10, 1, 6, 5, 7, 1, 2, 2, 3}

			for i, node := range heapValues {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}
		})

		t.Run("Test #2", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{10, 9, 5, 8, 2, 1, 3, 0, 4}

			heapValues := heap.GetElements()

			for i, node := range heapValues {
				if node.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], node.Value)
					break
				}
			}
		})

		t.Run("Test #3", func(t *testing.T) {
			values := []int{1, -5, 0, 3, 15, 11, 7, -2, 4}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{15, 4, 11, 3, 1, 0, 7, -5, -2}

			heapValues := heap.GetElements()

			for i, node := range heapValues {
				if node.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], node.Value)
					break
				}
			}
		})

		t.Run("Test #4", func(t *testing.T) {
			values := []int{10, 1, 0, -1, -1, 2, 3, 100, 55, 2, 20, 40}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{100, 55, 40, 10, 20, 3, 2, -1, 1, -1, 2, 0}

			heapValues := heap.GetElements()

			for i, node := range heapValues {
				if node.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], node.Value)
					break
				}
			}
		})
	})

	t.Run("Test Poll", func(t *testing.T) {
		t.Run("It should return the heap head", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			value := heap.Poll()
			expectedValue := 10

			if value.Value != expectedValue {
				t.Errorf("Expected %d, got %d\n", expectedValue, value.Value)
			}
		})

		t.Run("It should return the heap head #2", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			heap.Poll()
			value := heap.Poll()

			expectedValue := 9

			if value.Value != expectedValue {
				t.Errorf("Expected %d, got %d\n", expectedValue, value.Value)
			}

			expectedHeap := []int{8, 4, 5, 0, 2, 1, 3}
			heapElements := heap.GetElements()

			for i, node := range heapElements {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}
		})

		t.Run("It should return the heap head #3", func(t *testing.T) {
			values := []int{1, 1, 0, 3, 4}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			firstPoll := heap.Poll()
			expectedFirstPoll := 4

			if firstPoll.Value != expectedFirstPoll {
				t.Errorf("Expected %d, got %d\n", expectedFirstPoll, firstPoll.Value)
			}

			expectedHeap := []int{3, 1, 0, 1}
			heapElements := heap.GetElements()

			for i, heapValue := range heapElements {
				if heapValue.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], heapValue.Value)
					heap.Print()
					break
				}
			}

			secondPoll := heap.Poll()
			expectedSecondPoll := 3

			if secondPoll.Value != expectedSecondPoll {
				t.Errorf("Expected %d, got %d\n", expectedSecondPoll, secondPoll.Value)
			}

			expectedHeapAfterSecondPoll := []int{1, 1, 0}
			heapElements = heap.GetElements()

			for i, heapValue := range heapElements {
				if heapValue.Value != expectedHeapAfterSecondPoll[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeapAfterSecondPoll[i], heapValue.Value)
					heap.Print()
					break
				}
			}

			thirdPoll := heap.Poll()
			expectedThirdPoll := 1

			if thirdPoll.Value != expectedThirdPoll {
				t.Errorf("Expected %d, got %d\n", expectedThirdPoll, thirdPoll.Value)
			}

			expectedHeapAfterThirdPoll := []int{1, 0}
			heapElements = heap.GetElements()

			for i, heapValue := range heapElements {
				if heapValue.Value != expectedHeapAfterThirdPoll[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeapAfterThirdPoll[i], heapValue.Value)
					heap.Print()
					break
				}
			}
		})

		t.Run("It should return nil when the heap is empty", func(t *testing.T) {
			heap := priorityqueue.NewMaxHeap()
			got := heap.Poll()

			if got != nil {
				t.Errorf("Expected nil, got %d\n", got.Value)
			}
		})
	})

	t.Run("Test Remove", func(t *testing.T) {
		t.Run("It should return an error when the heap is empty", func(t *testing.T) {
			heap := priorityqueue.NewMaxHeap()
			err := heap.Remove(priorityqueue.NewNode(2))

			if err == nil {
				t.Fatal("Expected an error, got nil")
			}
		})

		t.Run("It should return an error when trying to remove an inexistent node", func(t *testing.T) {
			heap := priorityqueue.NewMaxHeap()
			heap.Insert(priorityqueue.NewNode(4))

			err := heap.Remove(priorityqueue.NewNode(2))

			if err == nil {
				t.Fatal("Expected an error, got nil")
			}
		})

		t.Run("Test #1", func(t *testing.T) {
			values := []int{3, 1, 0, 10, 11, 2, 4, -1}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(10))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
			}

			expectedHeap := []int{11, 3, 4, 1, -1, 0, 2}
			heapAfterRemove := heap.GetElements()

			for i, node := range heapAfterRemove {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}
		})

		t.Run("Test #2", func(t *testing.T) {
			values := []int{3, 1, 0, 10, 11, 2, 4, -1}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(10))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			err = heap.Remove(priorityqueue.NewNode(3))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap := []int{11, 2, 4, 1, -1, 0}
			heapAfterRemove := heap.GetElements()

			for i, node := range heapAfterRemove {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}
		})

		t.Run("Test #3", func(t *testing.T) {
			values := []int{3, 1, 0, -2, 10, 11, 2, 4, -1, 5, 7, 6}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(7))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap := []int{11, 5, 10, 3, 4, 6, 2, -2, -1, 1, 0}
			heapAfterFirstRemove := heap.GetElements()

			for i, node := range heapAfterFirstRemove {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}

			err = heap.Remove(priorityqueue.NewNode(3))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap = []int{11, 5, 10, 0, 4, 6, 2, -2, -1, 1}
			heapAfterSecondRemove := heap.GetElements()

			for i, node := range heapAfterSecondRemove {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}
		})
	})

	t.Run("Test Peek", func(t *testing.T) {
		t.Run("I should return the head without remove it", func(t *testing.T) {
			values := []int{3, 1, 0, -2, 10, 11, 2, 4, -1, 5, 7, 6}
			heap := priorityqueue.NewMaxHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			head := heap.Peak()

			if head == nil {
				t.Fatalf("Expected an element, got nil")
				return
			}

			heapElements := heap.GetElements()
			heapHead := heapElements[0]

			if head.Value != heapHead.Value {
				t.Errorf("Expected %d, got %d\n", heapHead.Value, head.Value)
			}
		})

		t.Run("It should return nil when the heap is empty", func(t *testing.T) {
			heap := priorityqueue.NewMaxHeap()
			head := heap.Peak()

			if head != nil {
				t.Errorf("Expected nil, got %d\n", head.Value)
			}
		})
	})
}
