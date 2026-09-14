package data_structure_test

import (
	"fmt"
	"libai/go/phase-two/data_structure"
	"testing"
)

func TestRing(t *testing.T) {
	window := data_structure.NewSlideWindow(5)
	for i := 0; i < 10; i++ {
		window.Push(float64(i))
		fmt.Printf("%d %f\n", i, window.Mean())
	}
}
