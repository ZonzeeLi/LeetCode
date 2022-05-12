## 剑指 Offer 17. 打印从1到最大的n位数

### 题目传送门

[点击这里](https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

### 解题思路

这道题不考虑大数情况的话就是一道很简单的题，只需要考虑n位数的最大数是多少即可，1位是9，2位是99，3位是999...以此类推，可以得出，n每多1，max要乘10然后加9。

### 代码

```go
func printNumbers(n int) []int {
    var res []int
    var max int
    // 1位是9，2位是99，3位是999，以此类推，可以得出，n每多1，max要乘10然后加9
    for n > 0 {
        max =max*10+9
    	n--
    }
    for i:=1;i<=max;i++{
    	res = append(res,i)
	}
	return res
}

```