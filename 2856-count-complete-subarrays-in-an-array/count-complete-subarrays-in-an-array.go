func countCompleteSubarrays(nums []int) int {
    uniqueNumsCnt := countUniqueNums(nums)
    lastOccurence := make(map[int]int)
    left := 0
    res := 0

    for right, num := range nums {
        lastOccurence[num] = right

        for left < lastOccurence[nums[left]] {
            left += 1
        }

        if len(lastOccurence) == uniqueNumsCnt {
            res += left + 1
        }
    }

    return res
}

func countUniqueNums(nums []int) int {
    uniqueNums := make(map[int]bool)
    
    for _, num := range nums {
        uniqueNums[num] = true
    }

    return len(uniqueNums)
}