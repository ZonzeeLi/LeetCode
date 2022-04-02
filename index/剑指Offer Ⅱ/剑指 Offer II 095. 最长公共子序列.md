## 剑指 Offer II 095. 最长公共子序列

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/qJnOS7/)

### 解题思路

这道题不难想到是使用动态规划，动态规划主要是三个点，1.dp数组的定义 2.边界条件的确定 3.状态转移方程 分析出这三点就不难写出动态规划的模板代码。

### 代码

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)

    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := range text1 {
        for j := range text2 {
            if text1[i] == text2[j] {
                dp[i+1][j+1] = dp[i][j] + 1
            }else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
