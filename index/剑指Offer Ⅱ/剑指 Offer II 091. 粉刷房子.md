## 剑指 Offer II 091. 粉刷房子

### 题目传送门

[点击这里](https://leetcode.cn/problems/JEj789/)

### 解题思路

这道题从初始的三个状态忘后进行计算，很明显的动态规划题，初始状态为三个颜色的状态，然后一直向后dp，由于相邻的房子颜色不能相同，所以我们能得到状态转移方程，`dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]`，`dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]`，`dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]`，根据规则，我们可以统一将其写成一条`dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]`，最后返回判断`dp[n-1]`三个状态中的最小值即可。又因为这道题`dp[i]`的状态一直和`dp[i-1]`有关系，所以可以优化降维，还可以使用滚动数组优化空间。

### 代码

```go
func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for j := 0; j < 3; j++ {
		dp[0][j] = costs[0][j]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
		// dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]
	}
	return min(min(dp[n-1][0], dp[n-1][1]), dp[n-1][2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

```go
func minCost(costs [][]int) int {
    dp := costs[0]
    for _, cost := range costs[1:] {
		// 滚动数组
        dpNew := make([]int, 3)
        for j, c := range cost {
            dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + c
        }
        dp = dpNew
    }
    return min(min(dp[0], dp[1]), dp[2])
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```
