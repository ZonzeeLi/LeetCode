package main

// 134. 加油站
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
func canCompleteCircuit(gas []int, cost []int) int {
	remain := make([]int, len(gas)*2)
	for k := range gas {
		remain[k] = gas[k] - cost[k]
	}
	// 找到油量剩余最小的位置即是起点，而且如果整体remain的和 >= 0，才说明有解
	count := 0
	minindex := 0
	mincount := 99999
	for k := range remain {
		count += remain[k]
		if mincount > count {
			mincount = count
			minindex = k
		}
	}
	if count < 0 {
		return -1
	}
	return (minindex + 1) % len(gas)
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := canCompleteCircuit(gas, cost)
	println(result) // 输出: 3
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	result = canCompleteCircuit(gas, cost)
	println(result) // 输出: -1
}
