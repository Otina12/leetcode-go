/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeInfo struct {
    node *TreeNode
    depth int
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    var resultNode NodeInfo = helper(root, 0)
    return resultNode.node
}

func helper(node *TreeNode, depth int) NodeInfo {
    if node == nil {
        return NodeInfo{nil, depth}
    }

    var leftChildInfo = helper(node.Left, depth + 1)
    var rightChildInfo = helper(node.Right, depth + 1)

    if leftChildInfo.depth == rightChildInfo.depth {
        return NodeInfo {node, leftChildInfo.depth}
    } else if leftChildInfo.depth > rightChildInfo.depth {
        return leftChildInfo
    } else {
        return rightChildInfo
    }
}