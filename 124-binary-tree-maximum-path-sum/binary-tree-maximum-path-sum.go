/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    var res int = root.Val
    var dfs func(*TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftMaxSum := max(0, dfs(node.Left))
        rightMaxSum := max(0, dfs(node.Right))

        res = max(res, node.Val + leftMaxSum + rightMaxSum)

        return node.Val + max(leftMaxSum, rightMaxSum)
    }

    dfs(root)
    return res
}