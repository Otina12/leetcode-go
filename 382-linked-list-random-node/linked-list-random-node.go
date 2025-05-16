/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
    listLength int
}


func Constructor(head *ListNode) Solution {
    len := 0
    copy := head

    for copy != nil {
        len += 1
        copy = copy.Next
    }

    return Solution{head, len}
}


func (this *Solution) GetRandom() int {
    nodeIdx := int(rand.Float64() * float64(this.listLength))
    copy := this.head

    for i := 0; i < nodeIdx; i++ {
        copy = copy.Next
    }

    return copy.Val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */