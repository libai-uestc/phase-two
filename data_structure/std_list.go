package data_structure

import (
	"container/list"
	"fmt"
)

func TraversList(lst *list.List) {
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next()
	}
	fmt.Println(head.Value)
}

func ReverseList(lst *list.List) {
	tail := lst.Back()
	for tail.Prev() != nil {
		fmt.Printf("%v ", tail.Value)
		tail = tail.Prev()
	}
	fmt.Println(tail.Value)
}
