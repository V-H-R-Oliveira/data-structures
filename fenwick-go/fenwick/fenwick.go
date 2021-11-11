package fenwick

import (
	"log"
)

// get lsb bit
func lsb(n int) int {
	return n - (n & -n)
}

// get next lsb bit
func nextLsb(n int) int {
	return n + (n & -n)
}

// Sum a given slice
func SumSubVector(elements ...int) int {
	if len(elements) == 1 {
		return elements[0]
	}

	sum := 0

	for _, element := range elements {
		sum += element
	}

	return sum
}

type FenwickTree []int

// Creates a new FenwickTree empty structure
func NewFenwick(size int) FenwickTree {
	return make(FenwickTree, size) // size -1 to start at 0
}

// Fills the structure with an array of values
func (fenwick FenwickTree) Fill(slice []int) {
	for i := 1; i < len(fenwick); i++ { // i := 1; i <= len(fenwick)
		lowIndex := lsb(i) // parent
		highIndex := lowIndex + (i - lowIndex)
		subSlice := slice[lowIndex:highIndex]
		fenwick[i] = SumSubVector(subSlice...) // i-1
	}
}

// Sum subSlice between Start and End indexes
func (fenwick FenwickTree) RangeSum(start, end int) int {
	fenwickLength := len(fenwick)

	if start < 1 || end >= fenwickLength {
		log.Fatalf("[RangeSum] Index out of range: 0 < %d < %d and 1 < %d < %d\n", start, fenwickLength, end, fenwickLength)
	}

	result := 0

	for parent := end; parent != 0; parent = lsb(parent) {
		result += fenwick[parent]
	}

	for parent := start - 1; parent != 0; parent = lsb(parent) {
		result -= fenwick[parent]
	}

	return result
}

// Update the Fenwick tree values given a specific range [a, b]
func (fenwick FenwickTree) Update(index, value int) {
	fenwickLength := len(fenwick)

	if index < 1 || index >= fenwickLength {
		log.Fatalf("[Update] Index out of range: 0 < %d < %d\n", index, fenwickLength)
	}

	for node := index; node <= len(fenwick); node = nextLsb(node) {
		fenwick[node] += value
	}
}
