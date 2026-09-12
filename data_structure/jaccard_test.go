// package data_structure_test

// import (
// 	"fmt"
// 	"libai/go/phase-two/data_structure"
// 	"testing"

// 	"slices"
// )

// func TestJaccardTimeConsuming(t *testing.T) {
// 	l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
// 	l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
// 	fmt.Println(data_structure.JaccardTimeConsuming(l1, l2))
// }

// func TestJaccardForSorted(t *testing.T) {
// 	l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
// 	l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
// 	slices.Sort(l1)
// 	slices.Sort(l2)
// 	fmt.Println(data_structure.JaccardForSorted(l1, l2))
// }

// // go test ./data_structure -v -run=^TestJaccardTimeConsuming$ -count=1
// // go test ./data_structure -v -run=^TestJaccardForSorted$ -count=1

package data_structure_test

import (
	"fmt"
	"libai/go/phase-two/data_structure"
	"slices"
	"testing"
)

func TestJaccardTimeConsuming(t *testing.T) {
	l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
	l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
	fmt.Println(data_structure.JaccardTimeConsuming(l1, l2))
	fmt.Println(l1)
	fmt.Println(l2)
	// fmt.Println(data_structure)
}

func TestJaccardForSorted(t *testing.T) {
	l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
	l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
	slices.Sort(l1)
	slices.Sort(l2)
	fmt.Println(data_structure.JaccardForSorted(l1, l2))
}

func TestHello(t *testing.T) {
	a := data_structure.Hello(10)
	fmt.Println(a)
}
