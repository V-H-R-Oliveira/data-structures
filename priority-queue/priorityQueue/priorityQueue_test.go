package priorityqueue_test

import (
	priorityqueue "priorityQueue/m/priorityQueue"
	"testing"
)

func TestMinHeap(t *testing.T) {
	t.Run("Test isEmpty", func(t *testing.T) {
		t.Run("Given an empty heap, it should return true", func(t *testing.T) {
			heap := priorityqueue.NewHeap()

			if !heap.IsEmpty() {
				t.Error("Expected an empty heap, got a filled heap")
			}
		})

		t.Run("Given a fulfilled heap, it should return false", func(t *testing.T) {
			values := []int{1, 5, 1, 8, 6, 2, 2, 13, 12, 11, 7, 2, 15, 3, 10}
			heap := priorityqueue.NewHeap()

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
			heap := priorityqueue.NewHeap()

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
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			heapValues := heap.GetElements()

			for i, value := range heapValues {
				if value.Value != values[i] {
					t.Errorf("Expected %d, got %d\n", values[i], value.Value)
					break
				}
			}
		})

		t.Run("Test #2", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{0, 1, 3, 5, 2, 4, 10, 9, 8}

			heapValues := heap.GetElements()

			for i, value := range heapValues {
				if value.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], value.Value)
					break
				}
			}
		})

		t.Run("Test #3", func(t *testing.T) {
			values := []int{1, -5, 0, 3, 15, 11, 7, -2, 4}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{-5, -2, 0, 1, 15, 11, 7, 3, 4}

			heapValues := heap.GetElements()

			for i, value := range heapValues {
				if value.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], value.Value)
					break
				}
			}
		})

		t.Run("Test #4", func(t *testing.T) {
			values := []int{10, 1, 0, -1, -1, 2, 3, 100, 55, 2, 20, 40}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			expected := []int{-1, -1, 1, 10, 0, 2, 3, 100, 55, 2, 20, 40}

			heapValues := heap.GetElements()

			for i, value := range heapValues {
				if value.Value != expected[i] {
					t.Errorf("Expected %d, got %d\n", expected[i], value.Value)
					break
				}
			}
		})
	})

	t.Run("Test Poll", func(t *testing.T) {
		t.Run("It should return the heap head", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			value := heap.Poll()
			expectedValue := 0

			if value.Value != expectedValue {
				t.Errorf("Expected %d, got %d\n", expectedValue, value.Value)
			}
		})

		t.Run("It should return the heap head #2", func(t *testing.T) {
			values := []int{4, 5, 1, 0, 2, 3, 10, 9, 8}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			heap.Poll()
			value := heap.Poll()

			expectedValue := 1

			if value.Value != expectedValue {
				t.Errorf("Expected %d, got %d\n", expectedValue, value.Value)
			}
		})

		t.Run("It should return the heap head #3", func(t *testing.T) {
			values := []int{1, 1, 0, 3, 4}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			firstPoll := heap.Poll()
			expectedFirstPoll := 0

			if firstPoll.Value != expectedFirstPoll {
				t.Errorf("Expected %d, got %d\n", expectedFirstPoll, firstPoll.Value)
			}

			expectedHeap := []int{1, 3, 1, 4}
			heapElements := heap.GetElements()

			for i, heapValue := range heapElements {
				if heapValue.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], heapValue.Value)
					heap.Print()
					break
				}
			}

			secondPoll := heap.Poll()
			expectedSecondPoll := 1

			if secondPoll.Value != expectedSecondPoll {
				t.Errorf("Expected %d, got %d\n", expectedSecondPoll, secondPoll.Value)
			}

			expectedHeapAfterSecondPoll := []int{1, 3, 4}
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

			expectedHeapAfterThirdPoll := []int{3, 4}
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
			heap := priorityqueue.NewHeap()
			got := heap.Poll()

			if got != nil {
				t.Errorf("Expected nil, got %d\n", got.Value)
			}
		})
	})

	t.Run("Test Remove", func(t *testing.T) {
		t.Run("It should return an error when the heap is empty", func(t *testing.T) {
			heap := priorityqueue.NewHeap()
			err := heap.Remove(priorityqueue.NewNode(2))

			if err == nil {
				t.Fatal("Expected an error, got nil")
			}
		})

		t.Run("It should return an error when trying to remove an inexistent node", func(t *testing.T) {
			heap := priorityqueue.NewHeap()
			heap.Insert(priorityqueue.NewNode(4))

			err := heap.Remove(priorityqueue.NewNode(2))

			if err == nil {
				t.Fatal("Expected an error, got nil")
			}
		})

		t.Run("Test #1", func(t *testing.T) {
			values := []int{3, 1, 0, 10, 11, 2, 4, -1}
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(0))

			if err != nil {
				t.Fatal("Failed due", err)
			}

			expectedHeap := []int{-1, 3, 1, 10, 11, 2, 4}
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
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(0))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			err = heap.Remove(priorityqueue.NewNode(2))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap := []int{-1, 3, 1, 10, 11, 4}
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
			heap := priorityqueue.NewHeap()

			for _, value := range values {
				heap.Insert(priorityqueue.NewNode(value))
			}

			err := heap.Remove(priorityqueue.NewNode(-2))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap := []int{-1, 0, 1, 3, 5, 6, 2, 4, 11, 10, 7}
			heapAfterFirstRemove := heap.GetElements()

			for i, node := range heapAfterFirstRemove {
				if node.Value != expectedHeap[i] {
					t.Errorf("Expected %d, got %d\n", expectedHeap[i], node.Value)
					break
				}
			}

			err = heap.Remove(priorityqueue.NewNode(0))

			if err != nil {
				t.Fatalf("Expected a nil error, got %v\n", err)
				return
			}

			expectedHeap = []int{-1, 3, 1, 4, 5, 6, 2, 7, 11, 10}
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
			heap := priorityqueue.NewHeap()

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
			heap := priorityqueue.NewHeap()
			head := heap.Peak()

			if head != nil {
				t.Errorf("Expected nil, got %d\n", head.Value)
			}
		})
	})
}
