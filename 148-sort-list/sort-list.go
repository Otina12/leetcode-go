/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow, fast := head, head
    temp := head

    for fast != nil && fast.Next != nil {
        temp = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    temp.Next = nil
    sortedLeft := sortList(head)
    sortedRight := sortList(slow)

    return mergeLists(sortedLeft, sortedRight)
}

func mergeLists(node1 *ListNode, node2 *ListNode) *ListNode {
    dummy := &ListNode{}
    mergedList := dummy

    for node1 != nil && node2 != nil {
        smallerVal := 0

        if node1.Val < node2.Val {
            smallerVal = node1.Val
            node1 = node1.Next
        } else {
            smallerVal = node2.Val
            node2 = node2.Next
        }

        dummy.Next = &ListNode{Val: smallerVal}
        dummy = dummy.Next
    }

    for node1 != nil {
        dummy.Next = &ListNode{Val: node1.Val}
        dummy = dummy.Next
        node1 = node1.Next
    }

    for node2 != nil {
        dummy.Next = &ListNode{Val: node2.Val}
        dummy = dummy.Next
        node2 = node2.Next
    }

    return mergedList.Next
}