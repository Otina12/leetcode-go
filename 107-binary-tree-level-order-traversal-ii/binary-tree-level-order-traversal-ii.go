/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    queue := []*TreeNode{root}
    levels := [][]int{}

    for len(queue) > 0 {
        size := len(queue)
        levels = append(levels, []int{})
        curLevelI := len(levels) - 1

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:] // O(1). In Go slices we trust!
            levels[curLevelI] = append(levels[curLevelI], node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    slices.Reverse(levels)
    return levels
}