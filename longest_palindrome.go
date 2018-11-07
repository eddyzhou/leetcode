package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end int
	for i, _ := range(s) {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		l := max(l1, l2)
		if l > (end-start) {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for ;left >= 0 && right < len(s) && s[left] == s[right]; {
		left--
		right++
	}
	return right-left-1
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
}