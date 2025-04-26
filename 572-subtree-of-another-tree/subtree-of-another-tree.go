/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    if root == nil || subRoot == nil {
        return false
    }

    return areSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func areSameTree(tree1 *TreeNode, tree2 *TreeNode) bool {
    if tree1 == nil && tree2 == nil {
        return true
    }
    if tree1 == nil || tree2 == nil {
        return false
    }

    return tree1.Val == tree2.Val && areSameTree(tree1.Left, tree2.Left) && areSameTree(tree1.Right, tree2.Right)
}