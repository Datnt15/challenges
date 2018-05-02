package main

type Node struct {
	Next  *Node
	Value interface{}
}

func (node *Node) Count() int {
	var count = 0
	for n := node; n != nil; n = n.Next {
		count++
	}
	return count
}

func Solution(list *Node, k int) *Node {
	var prev *Node = nil
	var node *Node = list
	var count = 0

	for node != nil && count <= k {
		prev = node
		node = node.Next
		count++
	}

	if prev != nil && node != nil {
		prev.Next = node.Next
		node.Next = nil
	}
	return node
}
