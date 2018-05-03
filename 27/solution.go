package main

import (
	"strings"
)

func Solution(input string) bool {

	const openBrackets = "({["
	const closeBrackets = ")}]"

	var queue []byte

	for i := 0; i < len(input); i++ {
		var r = rune(input[i])
		var b = strings.IndexRune(openBrackets, r)
		if b != -1 {
			queue = append(queue, openBrackets[b])
			continue
		}
		b = strings.IndexRune(closeBrackets, r)
		if b != -1 {
			pos := len(queue) - 1
			last := queue[pos]

			if last != openBrackets[b] {
				return false
			}
			queue = queue[:pos]
		}
	}
	return len(queue) == 0
}
