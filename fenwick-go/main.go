package main

import (
	fenwickLib "fenwick/m/fenwick"
	"fmt"
)

func main() {
	testFenwick := []int{3, 2, 4, 0, 42, 33, -1, -2, 4, 4}
	fenwick := fenwickLib.NewFenwick(len(testFenwick) + 1)

	fenwick.Fill(testFenwick)

	fmt.Println("Fenwick tree:", fenwick)

	fmt.Println(fenwick.RangeSum(3, 10), fenwickLib.SumSubVector(testFenwick[2:10]...))

	fmt.Println(
		fenwick.RangeSum(3, 5),
		fenwick.RangeSum(1, 10),
		fenwick.RangeSum(5, 8),
	)

	fenwick.Update(5, -2)

	fmt.Println(fenwick.RangeSum(3, 5))

	fenwick.Update(6, 7)

	fmt.Println(fenwick.RangeSum(4, 7))

	fenwick.RangeSum(4, 7)
}
