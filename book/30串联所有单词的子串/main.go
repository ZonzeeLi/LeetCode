package main

// 破题
func findSubstring(s string, words []string) []int {
	var res []int
	// 以示例"barfoothefoobarman", words = ["foo","bar"]为例
	ls, m, n := len(s), len(words), len(words[0])
	// 这里的i不要超过n，也不可以和m*n拼后超过ls的长度，后面的好理解，前面的目的是为了在后面的起始位置向前遍历不会有多余的n长度字符串影响，相当于初始遍历情况不要超过n开始。也就是官方解中的将s中的单词进行划分，切分为每个单词大小均为n，这样的划分方式有n中，即先删除前i个字符(i=[0~n-1])，对剩余进行划分，结尾剩余也删去。
	for i := 0; i < n && i+m*n <= ls; i++ {
		// 统计差异化的字符串
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		// 此时的differ情况分别在i为0、1、2时为 map[bar:1 foo:1]、map[arf:1 oot:1]、map[oth:1 rfo:1]
		// 对于words存在的word在differ中进行删除，因为出现过就可以弥补差异化
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				// 整体向后移动n格，计入新加入的n长度字符串，如果加入后在差异化map中value为0，说明补全了之前的一次差值，可以删除这个key
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				// 由于向后移动了，前面的n长度的单词要在差异化当中删掉
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				res = append(res, start)
			}
		}
	}
	return res
}
