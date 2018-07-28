package node

import "fmt"

type Node struct {
	Value interface{}
}

func New(val interface{}) *Node {
	return &Node{Value: val}
}

func (n *Node) Int() int {
	return n.Value.(int)
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}
