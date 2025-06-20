package main

import "strings"

// 68. 文本左对齐右对齐
// 给定一个单词数组 words 和一个长度 maxWidth，返回每行的文本左对齐或右对齐，
// 每行的长度为 maxWidth。每行的单词之间至少有一个空格，且不能超过 maxWidth。
// 如果当前行的单词数为 1，则该单词左对齐，行末填充剩余空格；
// 如果当前行的单词数大于 1，则平均分配空格，左边的单词之间多一个空格，
// 右边的单词之间少一个空格，行末填充剩余空格。
// 注意：
// 1. 最后一行的单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格；
// 2. 如果当前行的单词数为 1，则该单词左对齐，行末填充剩余空格；
// 3. 如果当前行的单词数大于 1，则平均分配空格，左边的单词之间多一个空格，
//    右边的单词之间少一个空格，行末填充剩余空格。
// 4. 如果当前行的单词数为 1，则该单词左对齐，行末填充剩余空格；

// 没什么意义的破题
// blank 返回长度为 n 的由空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}

func fullJustify(words []string, maxWidth int) (ans []string) {
	right, n := 0, len(words)
	for {
		left := right // 当前行的第一个单词在 words 的位置
		sumLen := 0   // 统计这一行单词长度之和
		// 循环确定当前行可以放多少单词，注意单词之间应至少有一个空格
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
		if right == n {
			s := strings.Join(words[left:], " ")
			ans = append(ans, s+blank(maxWidth-len(s)))
			return
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		// 当前行只有一个单词：该单词左对齐，在行末填充剩余空格
		if numWords == 1 {
			ans = append(ans, words[left]+blank(numSpaces))
			continue
		}

		// 当前行不只一个单词
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) // 拼接额外加一个空格的单词
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))  // 拼接其余单词
		ans = append(ans, s1+blank(avgSpaces)+s2)
	}
}
