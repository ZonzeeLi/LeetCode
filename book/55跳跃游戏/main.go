package main

// 55. 跳跃游戏
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
func canJump(nums []int) bool {
	a := 0
	for k, v := range nums {
		if k > a {
			break
		}
		if k+v >= a {
			a = k + v
		}
	}
	return a >= len(nums)-1
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := canJump(nums)
	println(result) // 输出: true
	nums = []int{3, 2, 1, 0, 4}
	result = canJump(nums)
	println(result) // 输出: false
}
