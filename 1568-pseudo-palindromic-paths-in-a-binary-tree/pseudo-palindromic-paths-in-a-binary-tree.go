/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pseudoPalindromicPaths (root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(node *TreeNode, curXor int) int {
    if node == nil {
        return 0
    }

    curXor = flipKthBit(curXor, node.Val)

    if node.Left == nil && node.Right == nil {
        if curXor & (curXor - 1) == 0 {
            return 1
        }
    }

    return dfs(node.Left, curXor) + dfs(node.Right, curXor)
}

func flipKthBit(number int, k int) int {
    var mask int = 1 << k
    return number ^ mask
}