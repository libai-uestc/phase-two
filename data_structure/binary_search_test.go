package data_structure_test

import (
	"libai/go/phase-two/data_structure"
	"math/rand"
	"slices"
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

func TestBinarySearch4Section(t *testing.T) {
	const L = 100
	for c := 0; c < 30; c++ {
		arr := make([]float64, 0, L)
		for i := 0; i < L+c; i++ {
			arr = append(arr, rand.Float64())
		}
		slices.Sort(arr)
		var target float64
		target = arr[0] - 1.0
		if data_structure.BinarySearch4Section(arr, target) != 0 {
			t.Fail()
		}
		target = arr[len(arr)-1] + 1.0
		if data_structure.BinarySearch4Section(arr, target) != len(arr) {
			t.Fail()
		}
		target = arr[0]
		if data_structure.BinarySearch4Section(arr, target) != 0 {
			t.Fail()
		}
		for i := 0; i < len(arr)-1; i++ {
			target = (arr[i] + arr[i+1]) / 2
			if data_structure.BinarySearch4Section(arr, target) != i+1 {
				t.Fail()
			}
			target = arr[i+1]
			if data_structure.BinarySearch4Section(arr, target) != i+1 {
				t.Fail()
			}
		}
	}
}

func TestBinarySearch_TableDriven(t *testing.T) {
	// 1. 定义测试表格（匿名结构体切片）
	tests := []struct {
		name     string // 测试用例的名字，方便报错时定位
		arr      []int  // 输入的数组
		target   int    // 查找的目标
		expected int    // 期望返回的索引
	}{
		// 在这里填入各种你想测试的情况（边界条件尤其重要）
		{"找到目标_在正中间", []int{1, 3, 5, 7, 9}, 5, 2},
		{"找到目标_在最左侧", []int{1, 3, 5, 7, 9}, 1, 0},
		{"找到目标_在最右侧", []int{1, 3, 5, 7, 9}, 9, 4},
		{"找不到目标_比所有数都小", []int{1, 3, 5, 7, 9}, 0, -1},
		{"找不到目标_比所有数都大", []int{1, 3, 5, 7, 9}, 10, -1},
		{"找不到目标_在数组中间缺失", []int{1, 3, 5, 7, 9}, 4, -1},
		{"空数组", []int{}, 5, -1},
		{"只有一个元素的数组_找到", []int{5}, 5, 0},
		{"只有一个元素的数组_找不到", []int{5}, 3, -1},
	}

	// 2. 遍历表格，执行测试
	for _, tt := range tests {
		// t.Run 会启动一个子测试
		t.Run(tt.name, func(t *testing.T) {
			got := data_structure.BinarySearch(tt.arr, tt.target)

			// 3. 对比实际结果和期望结果
			if got != tt.expected {
				// 使用 t.Errorf 打印清晰的错误信息
				t.Errorf("输入数组: %v, 查找目标: %d => 期望索引: %d, 实际得到: %d",
					tt.arr, tt.target, tt.expected, got)
			}
		})
	}
}
