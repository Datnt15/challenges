package queue

import "../node"

type Queue []*node.Node

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n *node.Node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *node.Node) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
