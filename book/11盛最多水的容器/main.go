package main

// 11. 盛最多水的容器
//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。

func maxArea(height []int) int {
	//双指针
	s, e := 0, len(height)-1
	var water int
	for s < e {
		if (e-s)*min(height[e], height[s]) > water {
			water = (e - s) * min(height[e], height[s])
		}
		if height[e] < height[s] {
			e--
		} else {
			s++
		}
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
