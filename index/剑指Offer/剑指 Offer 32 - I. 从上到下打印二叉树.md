## 剑指 Offer 32 - I. 从上到下打印二叉树

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/)

### 解题思路

这道题根据题意就是二叉树的层序遍历，层序遍历的方法就是bfs，这道题也是标准的bfs二叉树层序遍历应用的模板了。

### 代码

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    // 层序遍历bfs
    var res []int
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        q := stack[0]
        stack = stack[1:]
        if q.Left != nil {
            stack = append(stack, q.Left)
        }
        if q.Right != nil {
            stack = append(stack, q.Right)
        }
        res = append(res, q.Val)
    }
    return res
}
```