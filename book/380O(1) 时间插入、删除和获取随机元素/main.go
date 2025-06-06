package main

import "math/rand"

// 380. O(1) 时间插入、删除和获取随机元素
//实现RandomizedSet 类：
//
//RandomizedSet() 初始化 RandomizedSet 对象
//bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
//bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
//int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
//你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

type RandomizedSet struct {
	nums  []int
	index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.index[val]; has {
		return false
	}
	this.nums = append(this.nums, val)
	this.index[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, has := this.index[val]; !has {
		return false
	} else {
		// 首先将数组的最后一个元素替换到原val处
		v := this.nums[len(this.nums)-1]
		this.nums[i] = v
		this.nums = this.nums[:len(this.nums)-1]
		this.index[v] = i
		// 要注意这个delete要放在最后，因为可能存在原nums只有一个元素的情况，这时候如果插入0则会逻辑判断错误
		delete(this.index, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
