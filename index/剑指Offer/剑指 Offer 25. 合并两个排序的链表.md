## 剑指 Offer 25. 合并两个排序的链表

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)

### 解题思路

由于链表已经排序，我们只需要用双指针从两个链表的头节点开始比较然后生成一个新的链表即可，只需要注意将没有遍历完的链表直接添加到新链表的后面即可。

### 代码

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    newhead := new(ListNode)
    p, q, k := l1, l2, newhead
    for p != nil && q != nil {
        if p.Val < q.Val {
            k.Next = p
            p = p.Next
            k = k.Next
        }else {
            k.Next = q
            q = q.Next
            k = k.Next
        }
    }
    if p != nil {
        k.Next = p
    }else if q != nil {
        k.Next = q
    }
    return newhead.Next

}
```