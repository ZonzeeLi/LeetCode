## 剑指 Offer 12. 矩阵中的路径

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)

### 解题思路

这道题是典型的dfs+剪枝的配合运用题，其实也有回溯的原理，总之是经典的类别模板题了，dfs+剪枝的写法有很多，函数内闭包形式、记忆化形式、传参等等都有不同的写法。

### 代码

```go
func exist(board [][]byte, word string) bool {
   m,n := len(board), len(board[0])
   for i := 0; i < m; i++ {
       for j := 0; j < n; j++ {
        //如果在数组中找得到第一个数，就执行下一步，否则返回false
        if dfs(board, i, j, 0,word) {
                return true
            }
        }
    }
    return false   
}
func dfs(board [][]byte, i,j,k int, word string) bool {
    //如果找到最后一个数，则返回true,搜索成功
    if k == len(word) {
        return true
    }
    //i,j的约束条件
    if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
        return false
    }

    //进入DFS深度优先搜索
    //先把正在遍历的该值重新赋值，如果在该值的周围都搜索不到目标字符，则再把该值还原
    //如果在数组中找到第一个字符，则进入下一个字符的查找
    if board[i][j] == word[k] {
        temp := board[i][j]
        board[i][j] = ' '
        //下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
        //即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
        if dfs(board, i, j + 1, k + 1, word) || 
        dfs(board, i, j - 1, k + 1, word) || 
        dfs(board, i + 1, j, k + 1, word) ||
        dfs(board, i - 1, j, k + 1, word)  {
            return true
        }else {
        //没有找到目标字符，还原之前重新赋值的board[i][j]
            board[i][j] = temp
        }
    }
    return false
}
```