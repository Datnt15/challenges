package main

import "testing"

func TestSolution(t *testing.T) {
	const input = "([])[]({})"

	expected := true

	actual := Solution(input)

	if expected != actual {
		t.Fatal("expected ", expected, " actual ", actual)
	}
}

func TestSolutionFail1(t *testing.T) {
	const input = "([)]"

	expected := false

	actual := Solution(input)

	if expected != actual {
		t.Fatal("expected ", expected, " actual ", actual)
	}
}

func TestSolutionFail2(t *testing.T) {
	const input = "((()"

	expected := false

	actual := Solution(input)

	if expected != actual {
		t.Fatal("expected ", expected, " actual ", actual)
	}
}
