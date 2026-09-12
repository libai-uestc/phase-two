package data_structure

import (
	"cmp"
	"fmt"
)

type ListNode[T cmp.Ordered] struct {
	Value T
	Prev  *ListNode[T]
	Next  *ListNode[T]
}

type DoubleList[T cmp.Ordered] struct {
	Head   *ListNode[T]
	Tail   *ListNode[T]
	Length int
}

// 正序遍历
func (list *DoubleList[T]) Traverse() {
	curr := list.Head
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}

// 逆序遍历
func (list *DoubleList[T]) ReverseTraverse() {
	curr := list.Tail
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Prev
	}
	fmt.Println()
}

// 向尾部追加一个元素
func (list *DoubleList[T]) PushBack(x T) {
	node := &ListNode[T]{Value: x}
	tail := list.Tail
	if tail == nil { // 为空，如果只有一个元素，那么head和tail指向同一个唯一的元素
		list.Head = node
		list.Tail = node
	} else {
		tail.Next = node
		node.Prev = tail
		list.Tail = node
	}
	list.Length += 1
}

func (list *DoubleList[T]) PushFront(x T) {
	node := &ListNode[T]{Value: x}
	head := list.Head
	if head == nil {
		list.Head = node
		list.Tail = node
	} else {
		head.Prev = node
		node.Next = head
		list.Head = node
	}
	list.Length += 1
}

// 获取第idx个元素。O(N)
func (list *DoubleList[T]) Get(idx int) *ListNode[T] {
	if list.Length <= idx {
		return nil
	}
	curr := list.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	return curr
}

// 在n1后面添加一个元素x。O(N)
func (list *DoubleList[T]) InsertAfter(x T, n1 *ListNode[T]) {
	n2 := &ListNode[T]{Value: x}
	if n1.Next != nil {
		n3 := n1.Next
		n3.Prev = n2
		n2.Next = n3
	} else {
		list.Tail = n2
	}
	n1.Next = n2
	n2.Prev = n1
	list.Length += 1
}

// 在n3前面添加一个元素x。
func (list *DoubleList[T]) InsertBefore(x T, n3 *ListNode[T]) {
	n2 := &ListNode[T]{Value: x}
	if n3.Prev != nil {
		n1 := n3.Prev
		n1.Next = n2
		n2.Prev = n1
	} else {
		list.Head = n2
	}
	n3.Prev = n2
	n2.Next = n3
	list.Length += 1
}
