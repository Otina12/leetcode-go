func canPartitionKSubsets(nums []int, k int) bool {
    arrSum := 0

    for _, num := range nums {
        arrSum += num
    }

    if arrSum % k != 0 {
        return false
    }

    target := arrSum / k
    visited := make([]bool, len(nums))
    sort.Ints(nums)

    var backtrack func(int, int, int) bool
    backtrack = func(start int, curSum int, subsetCnt int) bool {
        if subsetCnt == k {
            return true
        }

        if curSum == target {
            return backtrack(0, 0, subsetCnt + 1)
        }

        for i := start; i < len(nums); i++ {
            newSum := curSum + nums[i]

            if visited[i] || newSum > target {
                continue
            }

            visited[i] = true
            if backtrack(i + 1, newSum, subsetCnt) {
                return true
            }
            visited[i] = false

            if curSum == 0 {
                break
            }
        }
        
        return false
    }

    return backtrack(0, 0, 0)
}