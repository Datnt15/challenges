package main

import (
	"../queue"
	"../stack"
)

func Solution(input *stack.Stack) {

	sz := input.Len()

	q := queue.New()
	s := stack.New()

	for i := 0; i < sz; i++ {
		if i < sz/2 {
			q.Push(input.Pop())
		} else {
			s.Push(input.Pop())
		}
	}

	for input.Len() < sz {
		if !s.Empty() {
			input.Push(s.Pop())
		}
		if !q.Empty() {
			input.Push(q.Pop())
		}
	}
}
