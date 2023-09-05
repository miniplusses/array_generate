package array_generate

import (
	"math/rand"
)

func ArrayGenerate(size int) []int {
	s := []int{}

	for x := 0; x < size; x++ {
		s = append(s, rand.Intn(10000))
	}

	return s
}
