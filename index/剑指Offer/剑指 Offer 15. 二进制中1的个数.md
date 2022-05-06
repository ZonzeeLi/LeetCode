## 剑指 Offer 15. 二进制中1的个数

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/)

### 解题思路

这道题很简单，而且这里用的是位运算的方法。可以是用`num & num-1`的方法去掉一个1，假设输入1010，减一的话就是1001，按位与的话就变成了1000，去掉了一个1，比如输入1011，减1的话就是1010，按位与是1010，去掉了一个1。

### 代码

```go
func hammingWeight(num uint32) int {
    var res int
    for num != 0 {
        num &= num-1
        res ++
    }
    return res
}
```