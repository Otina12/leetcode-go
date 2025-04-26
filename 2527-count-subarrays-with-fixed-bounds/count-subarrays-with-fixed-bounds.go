func countSubarrays(nums []int, minK int, maxK int) int64 {
    minKLastOccurence, maxKLastOccurence, badLastOccurence := -1, -1, -1
    var res int64 = 0

    for i, num := range nums {
        if num == minK {
            minKLastOccurence = i
        }
        if num == maxK {
            maxKLastOccurence = i
        }
        if num < minK || num > maxK {
            badLastOccurence = i
        }

        res += int64(max(0, min(minKLastOccurence, maxKLastOccurence) - badLastOccurence))
    }

    return res
}