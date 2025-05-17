/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type MemoKey struct {
    node *TreeNode
    robbedPrev bool
}

func rob(root *TreeNode) int {
    var memo = make(map[MemoKey]int)
    var dfs func(*TreeNode, bool) int

    dfs = func(node *TreeNode, robbedPrevious bool) int {
        if node == nil {
            return 0
        }

        memoKey := MemoKey{node, robbedPrevious}
        if val, exists := memo[memoKey]; exists {
            return val
        }

        res := dfs(node.Left, false) + dfs(node.Right, false)
        if !robbedPrevious {
            res = max(res, node.Val + dfs(node.Left, true) + dfs(node.Right, true))
        }
        
        memo[memoKey] = res
        return res
    }

    return dfs(root, false)
}