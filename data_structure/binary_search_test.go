package data_structure_test

import (
	"libai/go/phase-two/data_structure"
	"math/rand"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	const L = 100
LOOP:
	for i := 0; i < 100; i++ {
		arr := make([]int, L)
		for j := 0; j < L; j++ {
			arr[j] = rand.Intn(L)
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

		target := rand.Intn(L)
		index := data_structure.BinarySearch(arr, target)
		if index >= 0 {
			if arr[index] != target {
				t.Fail()
				break LOOP
			}
		} else {
			for _, ele := range arr {
				if ele == target {
					t.Fail()
					break LOOP
				}
			}
		}
	}
}
