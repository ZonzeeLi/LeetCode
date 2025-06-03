package main

// 121. 买卖股票的最佳时机
//给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

//func maxProfit(prices []int) int {
//	i, j := 0, 0
//	diff := 0
//	for ; j < len(prices); j++ {
//		if prices[j] < prices[i] {
//			i = j
//		} else {
//			diff = max(diff, prices[j]-prices[i])
//		}
//	}
//	return diff
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 比较好的解法
func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	println(result) // 输出: 5
}
