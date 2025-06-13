package main

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for j, c := range s0 {
		for _, s := range strs {
			if j == len(s) || s[j] != byte(c) {
				return s0[:j]
			}
		}
	}
	return s0
}
