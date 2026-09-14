package data_structure

import (
	"cmp"
	"errors"
	"slices"
)

type Heap[T cmp.Ordered] struct {
	arr []T
}

func NewHeap[T cmp.Ordered](arr []T) *Heap[T] {
	brr := slices.Clone(arr)
	return &Heap[T]{arr: brr}
}

func (heap *Heap[T]) downwardAdjust(parent int) {
	left := 2*parent + 1
	if left >= len(heap.arr) {
		return
	}

	minIndex := parent
	minValue := heap.arr[minIndex]
	if heap.arr[left] < minValue {
		minValue = heap.arr[left]
		minIndex = left
	}
	right := 2*parent + 2
	if right < len(heap.arr) {
		if heap.arr[right] < minValue {
			minIndex = right
		}
	}

	if minIndex != parent {
		heap.arr[minIndex], heap.arr[parent] = heap.arr[parent], heap.arr[minIndex]
		heap.downwardAdjust(minIndex)
	}
}

func (heap *Heap[T]) Build() {
	n := len(heap.arr)
	if n <= 1 {
		return
	}
	lastIndex := n / 2 * 2
	for i := lastIndex; i > 0; i -= 2 {
		right := i
		parent := (right - 1) / 2
		heap.downwardAdjust(parent)
	}
}

func (heap *Heap[T]) upwardAdjust(currIndex int) {
	if currIndex == 0 {
		return
	}
	parent := (currIndex - 1) / 2
	if heap.arr[currIndex] < heap.arr[parent] {
		heap.arr[currIndex], heap.arr[parent] = heap.arr[parent], heap.arr[currIndex]
		heap.upwardAdjust(parent)
	}
}

func (heap *Heap[T]) Push(ele T) {
	heap.arr = append(heap.arr, ele)
	heap.upwardAdjust(len(heap.arr) - 1)
}

func (heap *Heap[T]) Pop() (T, error) {
	if len(heap.arr) == 0 {
		var v T
		return v, errors.New("heap is empty")
	}
	return heap.arr[0], nil
}

func (heap *Heap[T]) Size() int {
	return len(heap.arr)
}

func (heap *Heap[T]) GetAll() []T {
	return heap.arr
}

func (heap *Heap[T]) ReplaceTop(ele T) {
	if len(heap.arr) == 0 {
		heap.Push(ele)
	} else {
		heap.arr[0] = ele
		heap.downwardAdjust(0)
	}
}
