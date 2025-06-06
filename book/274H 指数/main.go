package main

// 274. H 指数
// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
//
// 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
func hIndex(citations []int) int {
	//利用一个空间来记录引用次数出现的数量
	n := len(citations)
	i, j := 0, n
	for i < j {
		//二分这里的二分是将索引二分，所以是j+1
		mid := (i + j + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			i = mid
		} else {
			j = mid - 1
		}
	}
	return i
}

func main() {
	citations := []int{1, 3, 1}
	result := hIndex(citations)
	println(result) // 输出: 3
}
