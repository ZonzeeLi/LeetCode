## 剑指 Offer II 029. 排序的循环链表

### 题目传送门

[点击这里](https://leetcode.cn/problems/4ueAj6/)

### 解题思路

这道题的问题在于特殊情况的处理，当`aNode`为空，或者遍历的过程中进入了遍历循环，比如只有一个元素，或者元素相同。如果为空做单独处理，避免处理循环，我们可以`p := aNode`，然后`for p.Next != aNode`，在循环的其实条件避免让`p`遍历回来，之后就找头尾拐点或者是在某一中间位置可插入的点即可。

### 代码

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
    n := new(Node)
    n.Val = x
    if aNode == nil {
        n.Next = n
        return n
    }
    p := aNode
    // 这个避免一直循环，及时跳出
    for p.Next != aNode {
        if p.Val <= x && x <= p.Next.Val {
            break
        }
        if p.Val > p.Next.Val && (p.Next.Val >= x || p.Val <= x) {
            break
        }
        p = p.Next
    }
    n.Next = p.Next
    p.Next = n
    return aNode
}
```