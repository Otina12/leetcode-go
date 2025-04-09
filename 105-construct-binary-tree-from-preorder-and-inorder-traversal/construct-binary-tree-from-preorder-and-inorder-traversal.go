/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    var valueToIndex map[int]int = make(map[int]int)
    var buildSubtree func(int, int, int, int) *TreeNode
    var n = len(preorder)

    for i, val := range inorder {
        valueToIndex[val] = i
    }

    buildSubtree = func(pStart int, pEnd int, iStart int, iEnd int) *TreeNode {
        if pStart > pEnd {
            return nil
        }

        var subtreeRootVal = preorder[pStart]
        var subtreeRoot = &TreeNode { Val: subtreeRootVal }
        var inorderNodeIndex = valueToIndex[subtreeRootVal]
        var leftSubtreeSize = inorderNodeIndex - iStart

        subtreeRoot.Left = buildSubtree(pStart + 1, pStart + leftSubtreeSize, iStart, inorderNodeIndex - 1)
        subtreeRoot.Right = buildSubtree(pStart + leftSubtreeSize + 1, pEnd, inorderNodeIndex + 1, iEnd)

        return subtreeRoot
    }

    return buildSubtree(0, n - 1, 0, n - 1)
}