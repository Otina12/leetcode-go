/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Pair struct {
    Index int
    Value int
}

func nextLargerNodes(head *ListNode) []int {
    var monoStack []*Pair
    idxToAnswerMap := make(map[int]int)
    idx := 0

    for head != nil {
        pair := Pair{Index: idx, Value: head.Val}

        for len(monoStack) > 0 && pair.Value > monoStack[len(monoStack)-1].Value {
            poppedPair := monoStack[len(monoStack)-1]
            monoStack = monoStack[:len(monoStack)-1]
            idxToAnswerMap[poppedPair.Index] = pair.Value
        }

        monoStack = append(monoStack, &pair)
        head = head.Next
        idx += 1
    }

    res := make([]int, idx)

    for numI, answer := range idxToAnswerMap {
        res[numI] = answer
    }

    return res
}