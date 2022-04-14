## 剑指 Offer 24. 反转链表

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)

### 解题思路

反转链表是链表类型的基础题，这里介绍一下思路。`1->2->3->4->5->NULL`，我们可以先声明一个头节点`newhead`，然后优先存好`2`，将`1`和`2`断开，让`1->newhead`，然后`newhead`移动到`1`的位置，也就是头节点更新位置，这样子就成为了`1->nil`，同理，现在其实就是`1->nil 2->3->4->5->nil`两个链表了，依次进行下去。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    p := head 
    for p != nil {
        next := p.Next
        p.Next = pre
        pre = p
        p = next
    }
    return pre
}
```