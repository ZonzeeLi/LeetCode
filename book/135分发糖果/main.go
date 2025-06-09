package main

// 135. 分发糖果
// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
// 你需要按照以下要求，给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
func candy(ratings []int) int {
	count := len(ratings)
	for i := 0; i < len(ratings); i++ {
		start := i
		// start 作为最低点，向前校验下
		if i > 0 && ratings[i-1] < ratings[i] {
			start--
		}
		// 先找递增段
		for i < len(ratings)-1 && ratings[i] < ratings[i+1] {
			i++
		}
		// 找到了最高点
		top := i
		diffUp := (top - start) * (top - start - 1) / 2

		// 接着找递减段
		for i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			i++
		}

		// 找到了山的另一侧低点
		down := i
		diffdown := (down - top) * (down - top - 1) / 2
		count += diffUp + diffdown + max(top-start, down-top)
	}
	return count
}
