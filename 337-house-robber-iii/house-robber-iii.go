/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
    var dfs func(*TreeNode) (int, int)

    dfs = func(node *TreeNode) (int, int) { // returns (rob, notRob)
        if node == nil {
            return 0, 0
        }

        leftRob, leftNotRob := dfs(node.Left)
        rightRob, rightNotRob := dfs(node.Right)

        rob := node.Val + leftNotRob + rightNotRob
        notRob := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
        
        return rob, notRob
    }

    robRoot, notRobRoot := dfs(root)
    return max(robRoot, notRobRoot)
}