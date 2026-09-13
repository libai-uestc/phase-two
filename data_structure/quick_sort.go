package data_structure

import "cmp"

func Partition[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	pivot := 0
	i := 1
	j := len(arr) - 1
	for i < j {
		for ; i < j; j-- {
			if arr[j] < arr[pivot] {
				break
			}
		}
		for ; i < j; i++ {
			if arr[i] > arr[pivot] {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[i] < arr[pivot] {
		arr[pivot], arr[i] = arr[i], arr[pivot]
		pivot = i
	}
	if pivot > 0 {
		Partition(arr[:pivot])
	}
	if pivot < len(arr)-1 {
		Partition(arr[pivot+1:])
	}
}
