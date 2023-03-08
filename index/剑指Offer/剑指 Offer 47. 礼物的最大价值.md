## 剑指 Offer 47. 礼物的最大价值

### 题目传送门

[点击这里](https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/)

### 解题思路

这道题是典型的动态规划类型题，记录`dp[i][j]`，表示在位置`(i, j)`最多可以拿到的礼物价值，所以这道题的状态转移方程也是随着向右向下移动，比较清晰。注意到状态转移方程中，`dp[i][j]`只会从`dp[i-1][j]`和`dp[i][j-1]`转移，因此同一时刻只需要保存最后两行的状态即可。

### 代码

```go
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i, v := range grid {
		for j, g := range v {
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			dp[i][j] += g
		}
	}
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

```go
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 滚动数组的思想，如果只跟两个有关那么就压缩成两个。
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, v := range grid {
		pos := i % 2
		for j, g := range v {
			dp[pos][j] = 0
			if i > 0 {
				dp[pos][j] = max(dp[pos][j], dp[1-pos][j])
			}
			if j > 0 {
				dp[pos][j] = max(dp[pos][j], dp[pos][j-1])
			}
			dp[pos][j] += g
		}
	}
	return dp[(m-1)%2][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```