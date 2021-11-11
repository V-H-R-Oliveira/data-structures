package fenwick

import (
	"math/rand"
	"time"
)

// GenRandomSlice - Generates a random slice with n length
func GenRandomSlice(length int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		element := rand.Intn(length + 1)
		slice[i] = element
	}

	return slice
}
