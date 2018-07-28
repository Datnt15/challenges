package stack

import "../node"

type Stack []*node.Node

func New() *Stack {
	return &Stack{}
}

func (q *Stack) Push(n *node.Node) {
	*q = append(*q, n)
}

func (q *Stack) PushValue(val interface{}) {
	n := node.New(val)
	q.Push(n)
}

func (q *Stack) Pop() (n *node.Node) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}

func (q *Stack) Len() int {
	return len(*q)
}

func (q *Stack) Empty() bool {
	return len(*q) == 0
}
