package data_structure

import "cmp"

func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	begin := 0
	end := len(arr) - 1
	for begin <= end {
		middle := (begin + end) / 2
		if arr[middle] == target {
			return middle
		}
		if arr[middle] < target {
			begin = middle + 1
		}
		if arr[middle] > target {
			end = middle - 1
		}
	}
	return -1
}

func BinarySearch4Section[T cmp.Ordered](arr []T, target T) int {
	if len(arr) == 0 {
		return -1
	}
	begin := 0
	end := len(arr) - 1
	for {
		if arr[begin] > target {
			return begin
		}
		if arr[end] < target {
			return end + 1
		}
		middle := (begin + end) / 2
		if arr[middle] > target {
			end = middle - 1
		} else if arr[middle] < target {
			begin = middle + 1
		} else {
			return middle
		}
	}
}
