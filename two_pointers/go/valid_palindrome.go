//https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	test := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(test))
}

func isPalindrome(s string) bool {
	leftP, rightP := 0, len(s)-1

	for leftP < rightP {
		left, right := rune(s[leftP]), rune(s[rightP])

		if !isAlphaNumeric(left) {
			leftP++
			continue
		}

		if !isAlphaNumeric(right) {
			rightP--
			continue
		}

		if unicode.ToLower(left) != unicode.ToLower(right) {
			return false
		}

		leftP++
		rightP--
	}

	return true
}

func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}
