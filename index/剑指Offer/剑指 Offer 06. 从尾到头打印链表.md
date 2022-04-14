## 剑指 Offer 06. 从尾到头打印链表

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

### 解题思路

这道题可以把原链表的数拿出来然后反转数组，也可以先反转链表，也可以每次都从数组的头添加值。这里写一下反转链表的方法。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    // 可以先反转链表，也可以输出链表值再反转数组
    var res []int
    var cur *ListNode
    p := head
    for p != nil {
        n := p.Next
        p.Next = cur
        cur = p
        p = n
    }
    for cur != nil {
        res = append(res, cur.Val)
        cur = cur.Next
    }
    return res
}
```