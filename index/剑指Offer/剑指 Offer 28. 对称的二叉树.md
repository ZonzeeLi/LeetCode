## 剑指 Offer 28. 对称的二叉树

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)

### 解题思路

判断二叉树是否是镜像对称的，要判断的是左子树的左节点`l.left`和右子树的右节点`r.right`是否相同，左子树的右节点`l.right`和右子树的左节点`r.left`是否相同，同时要判断两个根节点是否相同，所以这也是一道递归问题。

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
func isSymmetric(root *TreeNode) bool {
    return recursion(root, root)
}

func recursion(l, r *TreeNode) bool {
    if l == nil && r == nil {
        return true
    }

    if l == nil || r == nil {
        return false
    }

    return l.Val == r.Val && recursion(l.Left, r.Right) && recursion(l.Right, r.Left)
}
```