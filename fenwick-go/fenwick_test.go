package main

import (
	fenwickLib "fenwick/m/fenwick"
	utils "fenwick/m/utils"
	"math/rand"
	"testing"
)

func TestFenwick(t *testing.T) {
	testSlice := utils.GenRandomSlice(1e6)

	lowerBound := int(5e5)
	upperBound := int(1e6)

	fen := fenwickLib.NewFenwick(len(testSlice) + 1)
	fen.Fill(testSlice)

	lower := 1 + rand.Intn(lowerBound+1)
	upper := lowerBound + rand.Intn(upperBound-lowerBound-1) + 1

	expected := fenwickLib.SumSubVector(testSlice[lower-1 : upper]...)
	got := fen.RangeSum(lower, upper)

	if expected != got {
		t.Errorf("Expect %d. Got %d\n", expected, got)
	}
}

func BenchmarkFenwickRangeSum(b *testing.B) {
	testSlice := utils.GenRandomSlice(1e6)

	lowerBound := int(5e5)
	upperBound := int(1e6)

	fen := fenwickLib.NewFenwick(len(testSlice) + 1)
	fen.Fill(testSlice)

	for i := 0; i < b.N; i++ {
		lower := 1 + rand.Intn(lowerBound+1)
		upper := lowerBound + rand.Intn(upperBound-lowerBound-1) + 1

		expected := fenwickLib.SumSubVector(testSlice[lower-1 : upper]...)
		got := fen.RangeSum(lower, upper)

		if expected != got {
			b.Fatalf("Expect %d. Got %d\n", expected, got)
		}
	}
}

func BenchmarkFenwickUpdate(b *testing.B) {
	testSlice := utils.GenRandomSlice(1e6)
	fen := fenwickLib.NewFenwick(len(testSlice) + 1)
	fen.Fill(testSlice)

	for i := 0; i < b.N; i++ {
		index, value := 1+rand.Intn(int(1e6)-1), rand.Int()
		fen.Update(index, value)
	}

}
