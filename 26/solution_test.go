package main

import (
	"math/rand"
	"testing"
)

func randomNode() *Node {
	return &Node{Value: rand.Int(), Next: nil}
}

func buildTestList(size int, k int) (*Node, *Node) {
	var start = randomNode()

	var node = start

	var value *Node = nil

	for i := 0; i < size; i++ {
		node.Next = randomNode()
		node = node.Next

		if i == k {
			value = node
		}
	}

	return start, value
}

func TestSolution(t *testing.T) {

	var k = 500

	const size = 1000

	var list, expected = buildTestList(size, k)

	var expectedSize = list.Count() - 1

	var actual = Solution(list, k)

	if actual == nil {
		t.Fatal("expected not nil solution")
	}

	if actual.Value != expected.Value {
		t.Fatal("expected ", expected.Value, " actual ", actual.Value)
	}

	var actualSize = list.Count()

	if actualSize != expectedSize {
		t.Fatal("expected list size ", expectedSize, " actual ", actualSize)
	}
}
