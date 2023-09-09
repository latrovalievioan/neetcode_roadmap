// https://leetcode.com/problems/group-anagrams/submissions/
// TODO: find a better solution
package main

import (
	"fmt"
)

func main() {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(test))
}

func groupAnagrams(strs []string) [][]string {
	checkedIndexes := make(map[int]bool)
	result := [][]string{}

	for i, v := range strs {
		_, ok := checkedIndexes[i]

		if ok {
			continue
		}

		currentGroup := []string{v}
		checkedIndexes[i] = true
		for j := i + 1; j < len(strs); j++ {
			_, ok := checkedIndexes[j]

			if ok {
				continue
			}

			if isAnagram(v, strs[j]) {
				checkedIndexes[j] = true
				currentGroup = append(currentGroup, strs[j])
			}
		}

		result = append(result, currentGroup)
	}

	return result
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := [123]int{}

	for i := range s {
		alphabet[int(s[i])]++
	}

	for i := range t {
		alphabet[int(t[i])]--
	}

	for i := 97; i < 123; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}
