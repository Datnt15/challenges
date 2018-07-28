package main

import "testing"
import "../stack"

func TestSolution(t *testing.T) {

	input := stack.New()

	input.PushValue(1)
	input.PushValue(2)
	input.PushValue(3)
	input.PushValue(4)
	input.PushValue(5)

	Solution(input)

	val := input.Pop()

	exp := 3

	if val.Int() != exp {
		t.Error("expected ", exp, " got ", val)
	}

	val = input.Pop()
	exp = 4

	if val.Int() != exp {
		t.Error("expected ", exp, " got ", val)
	}

	val = input.Pop()
	exp = 2

	if val.Int() != exp {
		t.Error("expected ", exp, " got ", val)
	}

	val = input.Pop()
	exp = 5

	if val.Int() != exp {
		t.Error("expected ", exp, " got ", val)
	}

	val = input.Pop()
	exp = 1

	if val.Int() != exp {
		t.Error("expected ", exp, " got ", val)
	}

	if input.Len() != 0 {
		t.Error("expected empty stack")
	}
}
