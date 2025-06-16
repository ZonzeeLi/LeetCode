package main

// 151. 反转字符串中的单词
//给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
//单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
//返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
//注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

//func reverseWords(s string) string {
//	var a string
//	var res string
//	for i := len(s) - 1; i >= 0; i-- {
//		if a == "" && s[i] == ' ' {
//			continue
//		}
//		if s[i] == ' ' {
//			res += a + " "
//			a = ""
//			continue
//		}
//		a = string(s[i]) + a
//	}
//	if a != "" {
//		res += a
//	}
//	if res[len(res)-1] == ' ' {
//		res = res[:len(res)-1] // 去掉最后的空格
//	}
//	return res
//}

func reverseWords(s string) string {
	// 去除首尾空格
	s = trimSpaces(s)
	i, j := len(s)-1, len(s)-1
	var res string
	for i >= 0 && j >= 0 {
		if s[i] == ' ' {
			res += s[i+1:j+1] + " "
			for s[i] == ' ' && i >= 0 {
				i--
			}
			j = i
			continue
		}
		i--
	}
	res = res + s[i+1:j+1]
	return res
}

func trimSpaces(s string) string {
	i := 0
	j := len(s) - 1
	for i < len(s) && s[i] == ' ' {
		i++
	}
	for j >= 0 && s[j] == ' ' {
		j--
	}
	if i > j {
		return ""
	}
	return s[i : j+1]
}

func main() {
	s := "the sky is blue"
	res := reverseWords(s)
	println(res) // Output: "blue is sky the"

	//s := "  hello world!  "
	//res := reverseWords(s)
	//println(res) // Output: "world! hello"
}
