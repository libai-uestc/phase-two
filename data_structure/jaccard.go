// package data_structure

// import "cmp"

// func JaccardTimeConsuming[T cmp.Ordered](collection1, collection2 []T) float64 {
// 	if len(collection1) == 0 || len(collection2) == 0 {
// 		return 0.0
// 	}
// 	map1 := make(map[T]struct{}, len(collection1))
// 	for _, ele := range collection1 {
// 		map1[ele] = struct{}{}
// 	}
// 	map2 := make(map[T]struct{}, len(collection2))
// 	for _, ele := range collection2 {
// 		map2[ele] = struct{}{}
// 	}
// 	intersection := 0 //交集的个数
// 	for key := range map1 {
// 		if _, exists := map2[key]; exists {
// 			intersection += 1
// 		}
// 	}
// 	return float64(intersection) / float64(len(collection1)+len(collection2)-intersection)
// }

// func JaccardForSorted[T cmp.Ordered](collection1, collection2 []T) float64 {
// 	if len(collection1) == 0 || len(collection2) == 0 {
// 		return 0.0
// 	}
// 	intersection := 0 //交集的个数
// 	for i, j := 0, 0; i < len(collection1) && j < len(collection2); {
// 		if collection1[i] == collection2[j] {
// 			intersection += 1
// 			i += 1
// 			j += 1
// 		} else if collection1[i] < collection2[j] {
// 			i += 1
// 		} else {
// 			j += 1
// 		}
// 	}
// 	return float64(intersection) / float64(len(collection1)+len(collection2)-intersection)
// }

package data_structure

import (
	"cmp"
	"fmt"
)

func JaccardTimeConsuming[T cmp.Ordered](a, b []T) float64 {
	map1 := make(map[T]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		map1[a[i]] = struct{}{}
	}

	map2 := make(map[T]struct{}, len(b))
	for j := 0; j < len(b); j++ {
		map2[b[j]] = struct{}{}
	}
	res := 0
	for k := range map1 {
		if _, exists := map2[k]; exists {
			res += 1
		}
	}
	return float64(res) / float64(len(a)+len(b)-res)
	// result := make([]T, 0)
	// for i := 0; i < len(b); i++ {
	// 	if _, ok := map1[b[i]]; ok {
	// 		result = append(result, b[i])
	// 	}
	// }
	// return result
}

// func JaccardForSorted[T cmp.Ordered](a, b []T) float64 {
// 	if len(a) == 0 || len(b) == 0 {
// 		return 0
// 	}
// 	var res []T
// 	i, j := 0, 0
// 	for {
// 		// i, j := 0, 0
// 		if a[i] == b[j] {
// 			res = append(res, a[i])
// 			i += 1 // 加上这行
// 			j += 1 // 加上这行
// 		}
// 		if a[i] > b[j] {
// 			j += 1
// 		}
// 		if a[i] < b[j] {
// 			i += 1
// 		}
// 		if i == len(a) || j == len(b) {
// 			break
// 		}
// 	}
// 	return float64(len(res)) / float64(len(a)+len(b)-len(res))
// }

func JaccardForSorted[T cmp.Ordered](a, b []T) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	var res []T
	i, j := 0, 0
	for {
		if a[i] == b[j] {
			res = append(res, a[i])
			i += 1
			j += 1
		} else if a[i] > b[j] { // Must use else if!
			j += 1
		} else if a[i] < b[j] { // Must use else if!
			i += 1
		}

		if i == len(a) || j == len(b) {
			break
		}
	}
	return float64(len(res)) / float64(len(a)+len(b)-len(res))
}

func Hello(a int) int {
	fmt.Println(a)
	return a + 1
}
