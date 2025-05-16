/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    var maxBinaryTreeBuilder func(int, int) *TreeNode

    maxBinaryTreeBuilder = func(left int, right int) *TreeNode {
        if left > right {
            return nil
        }

        maxVal, maxValIdx := -1, -1

        for i := left; i <= right; i++ {
            if nums[i] > maxVal {
                maxVal = nums[i]
                maxValIdx = i
            }
        }

        return &TreeNode{maxVal, maxBinaryTreeBuilder(left, maxValIdx - 1), maxBinaryTreeBuilder(maxValIdx + 1, right)}
    }

    return maxBinaryTreeBuilder(0, len(nums) - 1)
}