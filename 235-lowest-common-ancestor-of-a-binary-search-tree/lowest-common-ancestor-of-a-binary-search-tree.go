/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestorHelper(root, min(p.Val, q.Val), max(p.Val, q.Val))
}

func lowestCommonAncestorHelper(node *TreeNode, lessVal int, greaterVal int) *TreeNode {
    if node == nil || node.Val >= lessVal && node.Val <= greaterVal {
        return node
    }

    if node.Val > greaterVal {
        return lowestCommonAncestorHelper(node.Left, lessVal, greaterVal)
    } else {
        return lowestCommonAncestorHelper(node.Right, lessVal, greaterVal)
    }
}