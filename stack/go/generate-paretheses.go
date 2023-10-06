// https://leetcode.com/problems/generate-parentheses/
// TODO: solve again
package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := []string{}

	var generate func(n int, open int, closed int, slice string)
	generate = func(n int, open int, closed int, slice string) {
		if open == n && closed == n {
			fmt.Println(slice)
			result = append(result, slice)
		}

		if open < n {
			generate(n, open+1, closed, slice+"(")
		}

		if open > closed {
			generate(n, open, closed+1, slice+")")
		}
	}

	generate(n, 0, 0, "")

	return result
}
