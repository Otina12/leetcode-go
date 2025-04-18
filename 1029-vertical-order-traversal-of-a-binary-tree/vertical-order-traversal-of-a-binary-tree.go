/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeInfo struct {
    Node *TreeNode
    Vertical int
}

func verticalTraversal(root *TreeNode) [][]int {
    verticalValues := make(map[int][]int)
    curLevel := []NodeInfo { NodeInfo { root, 0 } }

    for len(curLevel) > 0 {
        nextLevel := []NodeInfo {}
        levelValues := make(map[int][]int)

        for _, nodeInfo := range curLevel {
            if nodeInfo.Node == nil {
                continue
            }

            vertical := nodeInfo.Vertical
            levelValues[vertical] = append(levelValues[vertical], nodeInfo.Node.Val)

            nextLevel = append(nextLevel,
                NodeInfo{Node: nodeInfo.Node.Left, Vertical: vertical - 1},
                NodeInfo{Node: nodeInfo.Node.Right, Vertical: vertical + 1})
        }

        for col, vals := range levelValues {
            sort.Ints(vals)
            verticalValues[col] = append(verticalValues[col], vals...)
        }

        curLevel = nextLevel
    }

    keys := make([]int, 0, len(verticalValues))
    for k := range verticalValues {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    res := make([][]int, 0, len(keys))

    for _, k := range keys {
        res = append(res, verticalValues[k])
    }

    return res
}