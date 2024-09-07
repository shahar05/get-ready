package utils

import (
	"log"
	"math/rand"
)

func Ptr[T any](v T) *T {
	return &v
}

// GetRandomInRange generates a random integer between min and max (inclusive)
func GetRandomInRange(min, max int) int {
	if min > max {
		log.Printf("Invalid range: min (%d) is greater than max (%d)\n", min, max)
		return -1 // -1 equal to error
	}
	return rand.Intn(max-min+1) + min
}
