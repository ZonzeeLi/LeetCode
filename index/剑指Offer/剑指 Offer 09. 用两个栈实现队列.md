## 剑指 Offer 09. 用两个栈实现队列

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

### 解题思路

一道简单题，用两个栈来实现队列，首先要对这两个数据结构有了解，栈是先入后出，队列是先入先出。已经说明使用两个栈来实现，所以结果定义就是两个`[]int`，一个用来实现`append`，一个用来实现`delete`，一个正序添加，一个倒序删除，倒序删除，就可以实现队列的先入先出特点。

### 代码

```go
type CQueue struct {
    stack1 []int // 用来实现append，正序存放
    stack2 []int // 用来实现delete，倒序存放
}


func Constructor() CQueue {
    return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
    this.stack1 = append(this.stack1, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.stack2) == 0 {
        if len(this.stack1) == 0 {
            return -1
        }
        // 将stack1的数加入到stack2中
        for len(this.stack1) > 0 {
            v := this.stack1[len(this.stack1)-1]
            this.stack1 = this.stack1[:len(this.stack1)-1]
            this.stack2 = append(this.stack2, v)
        }
    }
    v := this.stack2[len(this.stack2)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    return v
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
```