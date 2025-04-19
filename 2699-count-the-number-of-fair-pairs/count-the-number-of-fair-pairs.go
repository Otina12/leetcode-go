func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    n := len(nums)
    totalPairs := n * (n - 1) / 2
    unfairPairs := 0

    left, right := 0, n - 1
   
    for left < right {
        if nums[left] + nums[right] < lower {
            unfairPairs += right - left
            left += 1
        } else {
            right -= 1
        }
    }

    left, right = 0, n - 1

    for left < right {
        if nums[left] + nums[right] > upper {
            unfairPairs += right - left
            right -= 1
        } else {
            left += 1
        }
    }

    return int64(totalPairs - unfairPairs)
}