/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, lower int64, upper int64) bool {
    if node == nil {
        return true
    }

    curVal := int64(node.Val)

    if curVal <= lower || curVal >= upper {
        return false
    }

    return isValidBSTHelper(node.Left, lower, curVal) && isValidBSTHelper(node.Right, curVal, upper)
}