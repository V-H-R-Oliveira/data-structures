package main

import (
	"fmt"
	"testing"
)

func BenchmarkHashMapComparison(b *testing.B) {
	ranges := []struct {
		name  string
		size  int
		start uint64
	}{
		{"Size-64", 64, 16},        // 2^6, start 2^4
		{"Size-256", 256, 16},      // 2^8, start 2^4
		{"Size-1024", 1024, 16},    // 2^10, start 2^4
		{"Size-4096", 4096, 64},    // 2^12, start 2^6
		{"Size-16384", 16384, 128}, // 2^14, start 2^7
		{"Size-65536", 65536, 256}, // 2^16, start 2^8
	}

	for _, bm := range ranges {
		// Benchmark Custom HashMap - Put
		b.Run(bm.name+"/CustomHashMap/Put", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hm, _ := NewHashMap(bm.start)
				for j := 0; j < bm.size; j++ {
					hm.Put(keys[j], values[j])
				}
			}
		})

		// Benchmark Go Map - Put
		b.Run(bm.name+"/GoMap/Put", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := make(map[string]string, bm.start)
				for j := 0; j < bm.size; j++ {
					m[keys[j]] = values[j]
				}
			}
		})

		// Benchmark Custom HashMap - Get
		b.Run(bm.name+"/CustomHashMap/Get", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			hm, _ := NewHashMap(bm.start)
			for j := 0; j < bm.size; j++ {
				hm.Put(keys[j], values[j])
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < bm.size; j++ {
					hm.Get(keys[j])
				}
			}
		})

		// Benchmark Go Map - Get
		b.Run(bm.name+"/GoMap/Get", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			m := make(map[string]string, bm.start)
			for j := 0; j < bm.size; j++ {
				m[keys[j]] = values[j]
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < bm.size; j++ {
					_ = m[keys[j]]
				}
			}
		})

		// Benchmark Custom HashMap - PutAndGet
		b.Run(bm.name+"/CustomHashMap/PutAndGet", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hm, _ := NewHashMap(bm.start)
				for j := 0; j < bm.size; j++ {
					hm.Put(keys[j], values[j])
				}
				for j := 0; j < bm.size; j++ {
					hm.Get(keys[j])
				}
			}
		})

		// Benchmark Go Map - PutAndGet
		b.Run(bm.name+"/GoMap/PutAndGet", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := make(map[string]string, bm.start)
				for j := 0; j < bm.size; j++ {
					m[keys[j]] = values[j]
				}
				for j := 0; j < bm.size; j++ {
					_ = m[keys[j]]
				}
			}
		})

		// Benchmark Custom HashMap - Update
		b.Run(bm.name+"/CustomHashMap/Update", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			newValues := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
				newValues[i] = fmt.Sprintf("newvalue_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hm, _ := NewHashMap(bm.start)
				for j := 0; j < bm.size; j++ {
					hm.Put(keys[j], values[j])
				}
				for j := 0; j < bm.size; j++ {
					hm.Put(keys[j], newValues[j])
				}
			}
		})

		// Benchmark Go Map - Update
		b.Run(bm.name+"/GoMap/Update", func(b *testing.B) {
			keys := make([]string, bm.size)
			values := make([]string, bm.size)
			newValues := make([]string, bm.size)
			for i := 0; i < bm.size; i++ {
				keys[i] = fmt.Sprintf("key_%d", i)
				values[i] = fmt.Sprintf("value_%d", i)
				newValues[i] = fmt.Sprintf("newvalue_%d", i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := make(map[string]string, bm.start)
				for j := 0; j < bm.size; j++ {
					m[keys[j]] = values[j]
				}
				for j := 0; j < bm.size; j++ {
					m[keys[j]] = newValues[j]
				}
			}
		})
	}
}
