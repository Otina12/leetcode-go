// Keep track of previous prefix sums
// At each index i, we need to find largest j with j < i, where prefixSum[j] <= prefixSum[i] - k
// For example, prefixSum[i] = 9, k = 6. We need to find largest j at which prefixSum[j] is at most 3.

type Pair struct {
    Index int
    Value int
}

func shortestSubarray(nums []int, k int) int {
    monoQueue := []Pair{}
    curSum := 0
    res := len(nums) + 1

    for right, num := range nums {
        curSum += num

        if curSum >= k {
            res = min(res, right + 1)
        }

        for len(monoQueue) > 0 && curSum - monoQueue[0].Value >= k {
            res = min(res, right - monoQueue[0].Index)
            monoQueue = monoQueue[1:]
        }

        for len(monoQueue) > 0 && curSum <= monoQueue[len(monoQueue)-1].Value {
            monoQueue = monoQueue[:len(monoQueue)-1]
        }

        monoQueue = append(monoQueue, Pair{right, curSum})
    }

    if res == len(nums) + 1 {
        return -1
    }

    return res
}