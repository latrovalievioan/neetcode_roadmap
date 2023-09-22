// https://leetcode.com/problems/valid-parentheses
// TODO: solve again
package main

import (
	"fmt"
)

func main() {
	test := "]"
	fmt.Println(isValid(test))
}

func isValid(s string) bool {
	openingParenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	closingParenthesesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	sliceAsPseudoStack := []rune{}

	for _, v := range s {
		_, isOpening := openingParenthesesMap[v]

		if isOpening {
			sliceAsPseudoStack = append(sliceAsPseudoStack, v)
			continue
		}

		if len(sliceAsPseudoStack) <= 0 {
			return false
		}

		lastItem := sliceAsPseudoStack[len(sliceAsPseudoStack)-1]
		opening, _ := closingParenthesesMap[v]

		if lastItem != opening {
			return false
		}
		sliceAsPseudoStack = sliceAsPseudoStack[:len(sliceAsPseudoStack)-1]
	}

	return len(sliceAsPseudoStack) == 0
}
