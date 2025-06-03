package main

// 122. 买卖股票的最佳时机 II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	money := 0
	// dp[i][0] 表示第 i 天未持有股票（卖出）的利润，dp[i][1] 表示第 i 天持有股票（买入）的利润
	// 所以 dp[i][0] = dp[i-1][1] + prices[i], dp[i][1] = dp[i-1][0] - prices[i], 另外可以选择不买卖股票
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		money = max(money, dp[i][0])
		money = max(money, dp[i][1])
	}
	return money
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2, 3, 4, 5} // 测试用例
	result := maxProfit(prices)
	println(result) // 输出: 7
}
