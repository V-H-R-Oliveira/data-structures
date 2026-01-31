package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Benchmark data sizes
const (
	SmallSize  = 100
	MediumSize = 10000
	LargeSize  = 100000
)

const DefaultHashMapSize = 16

// Helper function to generate random keys
func generateKeys(n int) []string {
	keys := make([]string, n)

	for i := range n {
		keys[i] = fmt.Sprintf("key_%d_%d", i, rand.Intn(1000000))
	}
	return keys
}

// Helper function to generate random values
func generateValues(n int) []string {
	values := make([]string, n)

	for i := range n {
		values[i] = fmt.Sprintf("value_%d_%d", i, rand.Intn(1000000))
	}
	return values
}

// ============================================
// SMALL DATA BENCHMARKS (100 entries)
// ============================================

func BenchmarkCustomHashMap_Put_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	for b.Loop() {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := range SmallSize {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Put_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	for b.Loop() {
		m := make(map[string]string, 64)
		for j := range SmallSize {
			m[keys[j]] = values[j]
		}
	}
}

func BenchmarkCustomHashMap_Get_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)
	hm, _ := NewHashMap(DefaultHashMapSize)

	for j := range SmallSize {
		hm.Put(keys[j], values[j])
	}

	for b.Loop() {
		for j := range SmallSize {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_Get_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)
	m := make(map[string]string, 64)

	for j := range SmallSize {
		m[keys[j]] = values[j]
	}

	for b.Loop() {
		for j := range SmallSize {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_PutAndGet_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	for b.Loop() {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := range SmallSize {
			hm.Put(keys[j], values[j])
		}
		for j := range SmallSize {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_PutAndGet_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	for b.Loop() {
		m := make(map[string]string, 64)
		for j := range SmallSize {
			m[keys[j]] = values[j]
		}
		for j := range SmallSize {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_Update_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)
	newValues := generateValues(SmallSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < SmallSize; j++ {
			hm.Put(keys[j], values[j])
		}
		for j := 0; j < SmallSize; j++ {
			hm.Put(keys[j], newValues[j])
		}
	}
}

func BenchmarkGoMap_Update_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)
	newValues := generateValues(SmallSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string, 64)
		for j := 0; j < SmallSize; j++ {
			m[keys[j]] = values[j]
		}
		for j := 0; j < SmallSize; j++ {
			m[keys[j]] = newValues[j]
		}
	}
}

// ============================================
// MEDIUM DATA BENCHMARKS (10,000 entries)
// ============================================

func BenchmarkCustomHashMap_Put_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	for b.Loop() {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := range MediumSize {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Put_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	for b.Loop() {
		m := make(map[string]string)
		for j := range MediumSize {
			m[keys[j]] = values[j]
		}
	}
}

func BenchmarkCustomHashMap_Get_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	hm, _ := NewHashMap(DefaultHashMapSize)
	for j := 0; j < MediumSize; j++ {
		hm.Put(keys[j], values[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MediumSize; j++ {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_Get_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	m := make(map[string]string)
	for j := 0; j < MediumSize; j++ {
		m[keys[j]] = values[j]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MediumSize; j++ {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_PutAndGet_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < MediumSize; j++ {
			hm.Put(keys[j], values[j])
		}
		for j := 0; j < MediumSize; j++ {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_PutAndGet_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < MediumSize; j++ {
			m[keys[j]] = values[j]
		}
		for j := 0; j < MediumSize; j++ {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_Update_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)
	newValues := generateValues(MediumSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < MediumSize; j++ {
			hm.Put(keys[j], values[j])
		}
		for j := 0; j < MediumSize; j++ {
			hm.Put(keys[j], newValues[j])
		}
	}
}

func BenchmarkGoMap_Update_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)
	newValues := generateValues(MediumSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < MediumSize; j++ {
			m[keys[j]] = values[j]
		}
		for j := 0; j < MediumSize; j++ {
			m[keys[j]] = newValues[j]
		}
	}
}

// ============================================
// LARGE DATA BENCHMARKS (100,000 entries)
// ============================================

func BenchmarkCustomHashMap_Put_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Put_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < LargeSize; j++ {
			m[keys[j]] = values[j]
		}
	}
}

func BenchmarkCustomHashMap_Get_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	hm, _ := NewHashMap(DefaultHashMapSize)
	for j := 0; j < LargeSize; j++ {
		hm.Put(keys[j], values[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < LargeSize; j++ {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_Get_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	m := make(map[string]string)
	for j := 0; j < LargeSize; j++ {
		m[keys[j]] = values[j]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < LargeSize; j++ {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_PutAndGet_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], values[j])
		}
		for j := 0; j < LargeSize; j++ {
			hm.Get(keys[j])
		}
	}
}

func BenchmarkGoMap_PutAndGet_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < LargeSize; j++ {
			m[keys[j]] = values[j]
		}
		for j := 0; j < LargeSize; j++ {
			_ = m[keys[j]]
		}
	}
}

func BenchmarkCustomHashMap_Update_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)
	newValues := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], values[j])
		}
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], newValues[j])
		}
	}
}

func BenchmarkGoMap_Update_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)
	newValues := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < LargeSize; j++ {
			m[keys[j]] = values[j]
		}
		for j := 0; j < LargeSize; j++ {
			m[keys[j]] = newValues[j]
		}
	}
}

// ============================================
// REHASHING SPECIFIC BENCHMARKS
// ============================================

func BenchmarkCustomHashMap_Rehashing_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Start with very small size to force rehashing
		hm, _ := NewHashMap(8)
		for j := 0; j < SmallSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkCustomHashMap_Rehashing_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Start with very small size to force multiple rehashings
		hm, _ := NewHashMap(8)
		for j := 0; j < MediumSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkCustomHashMap_Rehashing_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Start with very small size to force multiple rehashings
		hm, _ := NewHashMap(8)
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

// ============================================
// COLLISION HEAVY BENCHMARKS
// ============================================

func BenchmarkCustomHashMap_HighCollision_Small(b *testing.B) {
	// Use sequential keys to potentially cause more collisions
	keys := make([]string, SmallSize)
	values := make([]string, SmallSize)
	for i := 0; i < SmallSize; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
		values[i] = fmt.Sprintf("value%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(16) // Small initial size to increase collisions
		for j := 0; j < SmallSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_HighCollision_Small(b *testing.B) {
	// Use sequential keys to potentially cause more collisions
	keys := make([]string, SmallSize)
	values := make([]string, SmallSize)
	for i := 0; i < SmallSize; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
		values[i] = fmt.Sprintf("value%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string, 16)
		for j := 0; j < SmallSize; j++ {
			m[keys[j]] = values[j]
		}
	}
}

// ============================================
// MEMORY ALLOCATION BENCHMARKS
// ============================================

func BenchmarkCustomHashMap_Memory_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < SmallSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Memory_Small(b *testing.B) {
	keys := generateKeys(SmallSize)
	values := generateValues(SmallSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string, 64)
		for j := 0; j < SmallSize; j++ {
			m[keys[j]] = values[j]
		}
	}
}

func BenchmarkCustomHashMap_Memory_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < MediumSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Memory_Medium(b *testing.B) {
	keys := generateKeys(MediumSize)
	values := generateValues(MediumSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < MediumSize; j++ {
			m[keys[j]] = values[j]
		}
	}
}

func BenchmarkCustomHashMap_Memory_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm, _ := NewHashMap(DefaultHashMapSize)
		for j := 0; j < LargeSize; j++ {
			hm.Put(keys[j], values[j])
		}
	}
}

func BenchmarkGoMap_Memory_Large(b *testing.B) {
	keys := generateKeys(LargeSize)
	values := generateValues(LargeSize)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		for j := 0; j < LargeSize; j++ {
			m[keys[j]] = values[j]
		}
	}
}
