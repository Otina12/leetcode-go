/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    var prev = &ListNode{}
    prev.Next = head
    var cur = head
    var dummy = prev

    for cur != nil && cur.Next != nil {
        var temp = cur.Next.Next
        prev.Next = cur.Next
        cur.Next.Next = cur
        cur.Next = temp
        prev = cur
        cur = cur.Next
    }

    return dummy.Next
}