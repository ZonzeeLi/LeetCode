package main

// 3. 无重复字符的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

func lengthOfLongestSubstring(s string) int {
	exist := make(map[rune]struct{})
	var res int
	start := 0
	for i := 0; i < len(s); i++ {
		// 如果遇到相同的，则左边界向后移动，并且删除当前在map中的key，直到删到当前的值
		if _, ok := exist[rune(s[i])]; ok {
			for start < len(s) {
				if s[start] == s[i] {
					start++
					break
				}
				delete(exist, rune(s[start]))
				start++
			}
		} else {
			exist[rune(s[i])] = struct{}{}
			res = max(res, i-start+1)
		}
	}
	return res
}

func main() {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	println(result) // Output: 3
}
