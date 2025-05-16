/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    var visited = make(map[int]struct{})
    var inorderTraversalHelper func(*TreeNode) bool

    inorderTraversalHelper = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        if _, exists := visited[k - node.Val]; exists {
            return true
        }

        visited[node.Val] = struct{}{}

        return inorderTraversalHelper(node.Left) || inorderTraversalHelper(node.Right)
    }

    return inorderTraversalHelper(root)
}