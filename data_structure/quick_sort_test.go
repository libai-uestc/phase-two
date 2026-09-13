package data_structure_test

import (
	"fmt"
	"libai/go/phase-two/data_structure"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		slice := make([]int, 20)
		for j := 0; j < 20; j++ {
			slice[j] = rand.Intn(100)
		}
		data_structure.Partition(slice)
		for j := 1; j < len(slice); j++ {
			if slice[j] < slice[j-1] {
				t.Fail()
			}
		}
	}
}

func TestQuickSort2(t *testing.T) {
	arr := []int{4, 3, 6, 1, 27}
	data_structure.Partition(arr)
	fmt.Println(arr)
}
