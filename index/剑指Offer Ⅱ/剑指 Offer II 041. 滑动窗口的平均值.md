## 剑指 Offer II 041. 滑动窗口的平均值

### 题目传送门

[点击这里](https://leetcode.cn/problems/qIsx9U/)

### 解题思路

题目的意思很简单，这道题也不难，其实说是实现滑动窗口，最终是用队列来完成，整体思路不难理解。

### 代码

```go
type MovingAverage struct {
    size int
    queue []int
    sum int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{size: size}
}


func (this *MovingAverage) Next(val int) float64 {
    if len(this.queue) == this.size {
        // 满了
        this.sum -= this.queue[0]
        this.queue = this.queue[1:]
    }
    this.sum += val
    this.queue = append(this.queue, val)
    return float64(this.sum) / float64(len(this.queue))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */


```