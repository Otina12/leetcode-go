/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec {}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return serializeHelper(root, 0)
}

func serializeHelper(root *TreeNode, depth int) string {
	if root == nil {
		return ""
	}

	depthSeparator := fmt.Sprintf("(%d)", depth)
	val := strconv.Itoa(root.Val)
	leftSubtreeString := serializeHelper(root.Left, depth + 1)
	rightSubtreeString := serializeHelper(root.Right, depth + 1)

	return val + depthSeparator + leftSubtreeString + depthSeparator + rightSubtreeString
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    return deserializeHelper(data, 0)
}

func deserializeHelper(data string, depth int) *TreeNode {
    if len(data) == 0 {
        return nil
    }

    depthSeparator := fmt.Sprintf("(%d)", depth)
    parts := strings.Split(data, depthSeparator)

    parentString, leftSubtreeString, rightSubtreeString := parts[0], parts[1], parts[2]
    parentVal, _ := strconv.Atoi(parentString)

    parent := &TreeNode { Val: parentVal }
    parent.Left = deserializeHelper(leftSubtreeString, depth + 1)
    parent.Right = deserializeHelper(rightSubtreeString, depth + 1)

    return parent
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */