/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    lenA, lenB := getLength(headA), getLength(headB)
    lenDiff := abs(lenA - lenB)
    smaller, larger := headA, headB

    if lenA > lenB {
        temp := smaller
        smaller = larger
        larger = temp
    }

    for i := 0; i < lenDiff; i++ {
        larger = larger.Next
    }

    for smaller != nil {
        if smaller == larger {
            return smaller
        }
        smaller = smaller.Next
        larger = larger.Next
    }

    return nil
}

func getLength(node *ListNode) int {
    if node == nil {
        return 0
    }

    return 1 + getLength(node.Next)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}