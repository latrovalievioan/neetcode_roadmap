package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	split1 := strings.Split(s, "")
	split2 := strings.Split(t, "")

	sort.Strings(split1)
	sort.Strings(split2)

	join1 := strings.Join(split1, "")
	join2 := strings.Join(split2, "")

	return join1 == join2
}

// solution 2
func isAnagram2(s string, t string) bool {
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
