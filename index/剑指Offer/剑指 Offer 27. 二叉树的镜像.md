## 剑指 Offer 27. 二叉树的镜像

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

### 解题思路

二叉树的遍历处理题，这种题不难想到用递归来做，我们要优先对深层处理，然后再逐级向上，所以先递归，再处理，先交换最深层的左右节点，然后向上回滚。

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
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := mirrorTree(root.Left)
    right := mirrorTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

```