/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    clones := make(map[int]*Node)
    clones[node.Val] = &Node { Val: node.Val }
    
    queue := []*Node { node }

    for len(queue) != 0 {
        curNode := queue[len(queue)-1]
        queue = queue[:len(queue)-1]
        curClone := clones[curNode.Val]

        for _, nei := range curNode.Neighbors {
            if _, exists := clones[nei.Val]; !exists {
                clones[nei.Val] = &Node { Val: nei.Val }
                queue = append(queue, nei)
            }
            
            clones[curClone.Val].Neighbors = append(clones[curClone.Val].Neighbors, clones[nei.Val])   
        }
    }

    return clones[node.Val]
}