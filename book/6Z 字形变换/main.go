package main

import "strings"

// 6. Z 字形变换
//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}
	var arr [][]string
	for i := 0; i < numRows; i++ {
		arr = append(arr, []string{})
	}
	j := 1
	head := true
	for i := 0; i < len(s); i++ {
		arr[j-1] = append(arr[j-1], string(s[i]))
		if head {
			j++
		} else {
			j--
		}
		if j == numRows || j == 1 {
			head = !head
		}
	}
	var res string
	for i := 0; i < len(arr); i++ {
		res += strings.Join(arr[i], "")
	}
	return res
}

func main() {
	//s := "PAYPALISHIRING"
	//numRows := 3
	s := "AB"
	numRows := 1
	result := convert(s, numRows)
	println(result) // Output: "PAHNAPLSIIGYIR"
}
