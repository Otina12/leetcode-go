/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    var dfs func(*TreeNode, int, int)
    maxDiff := 0

    dfs = func(node *TreeNode, maxVal int, minVal int) {
        if node == nil {
            return
        }

        maxDiff = max(maxDiff, abs(maxVal - node.Val), abs(minVal - node.Val))
        newMaxVal := max(maxVal, node.Val)
        newMinVal := min(minVal, node.Val)

        dfs(node.Left, newMaxVal, newMinVal)
        dfs(node.Right, newMaxVal, newMinVal)
    }

    dfs(root, root.Val, root.Val)
    return maxDiff
}

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}