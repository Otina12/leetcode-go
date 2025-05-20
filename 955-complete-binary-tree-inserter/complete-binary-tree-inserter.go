/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type CBTInserter struct {
    root *TreeNode
    candidates []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    queue := []*TreeNode{root}
    candidates := []*TreeNode{}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }

        if node.Left == nil || node.Right == nil {
            candidates = append(candidates, node)
        }
    }

    return CBTInserter{root, candidates}
}

func (this *CBTInserter) Insert(val int) int {
    nodeToInsert := &TreeNode{Val: val}
    parent := this.candidates[0]

    if parent.Left == nil {
        parent.Left = nodeToInsert
    } else if parent.Right == nil {
        parent.Right = nodeToInsert
        this.candidates = this.candidates[1:]
    }

    this.candidates = append(this.candidates, nodeToInsert)
    return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
    return this.root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */