package data_structure_test

import (
	"container/heap"
	"fmt"
	"libai/go/phase-two/data_structure"
	"testing"
)

func TestHeap(t *testing.T) {
	h := data_structure.NewHeap([]int{10, 15, 49, 20, 30, 62})
	fmt.Println(h.Size())
	h.Build()
	h.Push(5)
	for h.Size() > 0 {
		top, _ := h.Pop()
		fmt.Println(top)
	}
}

func TestStdHeap(t *testing.T) {
	pq := make(data_structure.PriorityQueue[int], 0, 10)
	pq.Push(&data_structure.Item[int]{Info: "A", Value: 3})
	pq.Push(&data_structure.Item[int]{Info: "B", Value: 2})
	pq.Push(&data_structure.Item[int]{Info: "C", Value: 4})
	heap.Init(&pq)
	heap.Push(&pq, &data_structure.Item[int]{Info: "D", Value: 6})
	for pq.Len() > 0 {
		fmt.Println(heap.Pop(&pq))
	}

}
