## 剑指 Offer 05. 替换空格

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

### 解题思路

一道简单题，go语言有内置的strings包下的替换函数，也可以使用遍历的做法，这里提一点，最好不要直接`var res string`，然后`res += "%20"`这样的写法，这样的写法不如`strings.Builder{}`，因为就如官方对`strings.Builder{}`的解释一样，`A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.`

```go
func replaceSpace(s string) string {
	res := strings.Builder{}
	for _, v := range s {
		if v == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteRune(v)
		}
	}
	return res.String()
}
```