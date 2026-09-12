package data_structure_test

import (
	"libai/go/phase-two/data_structure"
	"testing"
)

var binaryTree *data_structure.BNode

func init() {
	binaryTree = &data_structure.BNode{Value: 5}
	n15 := &data_structure.BNode{Value: 15}
	n10 := &data_structure.BNode{Value: 10}
	n20 := &data_structure.BNode{Value: 20}
	n30 := &data_structure.BNode{Value: 30}
	n62 := &data_structure.BNode{Value: 62}
	n49 := &data_structure.BNode{Value: 49}
	binaryTree.LeftChild = n15
	binaryTree.RightChild = n10
	n15.LeftChild = n20
	n15.RightChild = n30
	n10.LeftChild = n62
	n10.RightChild = n49

}

func TestPreOrder(t *testing.T) {
	binaryTree.PreOrder()
}

func TestPostOrder(t *testing.T) {
	binaryTree.PostOrder()
}
func TestMiddleOrder(t *testing.T) {
	binaryTree.MiddleOrder()
}
