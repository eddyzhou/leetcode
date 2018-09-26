package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	var start, max int
	for i, c := range(s) {
		if idx, ok := m[string(c)]; ok {
			if idx + 1 > start {
				start = idx + 1
			}
		}
		diff := i - start + 1
		if diff > max {
			max = diff
		}
		m[string(c)] = i
	}
	return max
}

//func main() {
//	s1 := "abcabcbb"
//	fmt.Println(lengthOfLongestSubstring(s1))
//	s2 := "bbbbb"
//	fmt.Println(lengthOfLongestSubstring(s2))
//	s3 := "pwwkew"
//	fmt.Println(lengthOfLongestSubstring(s3))
//	s4 := "abcdefg"
//	fmt.Println(lengthOfLongestSubstring(s4))
//}
