func resultsArray(nums []int, k int) []int {
    n := len(nums)
    lastInvalidIdx := 0
    var results []int
    
    for i := 0; i < n; i++ {
        if i > 0 && nums[i] != nums[i-1] + 1 {
            lastInvalidIdx = i
        }

        if i < k - 1 {
            continue
        }
        if i - lastInvalidIdx + 1 >= k {
            results = append(results, nums[i])
        } else {
            results = append(results, -1)
        }
    }

    return results
}