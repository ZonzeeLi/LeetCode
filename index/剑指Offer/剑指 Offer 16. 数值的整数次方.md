## 剑指 Offer 16. 数值的整数次方

### 题目传送门

[点击这里](https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

### 解题思路

这道题题干说明不需要考虑大数问题，所以只需要考虑时间复杂度和浮点数输入，这里可以使用递归和迭代，递归的相对于迭代会产生更多的栈来处理函数，浪费空间，不过也可以做，迭代和递归的思路一样，都是将n次幂，每次除以2，让结果计算`x*x`，最后剩余1，再乘一次。

### 代码

```go
func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickMul(x, n)
    }
    return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
    var res = 1.0
    product := x 
    for n > 0 {
        if n % 2 == 1 {
            // 奇数要多乘一次，最终的结果会统计到1
            res *= product
        }
        // 剩下的以平方计算
        n /= 2
        product *= product // 最后n剩下1的时候，该计算没有意义。
    }
    return res
}
```