package main

// 45. 跳跃游戏 II
//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
//每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 复杂度较高的方法
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 100000
	}
	dp[0] = 0
	for k, v := range nums {
		a := v
		for i := 0; i <= a && k+i < len(nums); i++ {
			dp[k+i] = min(dp[k]+1, dp[k+i])

		}
	}
	return dp[len(nums)-1]
}

//func jump(nums []int) (ans int) {
//	curRight := 0  // 已建造的桥的右端点
//	nextRight := 0 // 下一座桥的右端点的最大值
//	for i, num := range nums[:len(nums)-1] {
//		// 遍历的过程中，记录下一座桥的最远点
//		nextRight = max(nextRight, i+num)
//		if i == curRight { // 无路可走，必须建桥
//			curRight = nextRight // 建桥后，最远可以到达 next_right
//			ans++
//		}
//	}
//	return
//}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := jump(nums)
	println(result) // 输出: 2
}
