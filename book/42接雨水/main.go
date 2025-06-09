package main

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// https://www.bilibili.com/video/BV1Qg411q7ia/?vd_source=20b65b37d58ba62c59075224d1f2f64a
//func trap(height []int) int {
//	n := len(height)
//	// 当前i和i之前的最大值
//	preMax := make([]int, n)
//	sufMax := make([]int, n)
//	preMax[0] = height[0]
//	for i := 1; i < n; i++ {
//		preMax[i] = max(preMax[i-1], height[i])
//	}
//	sufMax[n-1] = height[n-1]
//	for i := n - 2; i >= 0; i-- {
//		sufMax[i] = max(sufMax[i+1], height[i])
//	}
//
//	count := 0
//	for i := 0; i < n; i++ {
//		d := min(preMax[i], sufMax[i]) - height[i]
//		if d > 0 {
//			count += d
//		}
//	}
//	return count
//}

func trap(height []int) int {
	count := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax < rightMax {
			count += max(0, leftMax-height[left])
			left++
			leftMax = max(leftMax, height[left])
		} else {
			count += max(0, rightMax-height[right])
			right--
			rightMax = max(rightMax, height[right])
		}
	}
	return count
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap(height)
	println(result) // 输出: 6
}
