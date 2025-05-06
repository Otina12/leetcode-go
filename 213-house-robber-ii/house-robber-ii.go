func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    return max(robHelper(nums, 1, len(nums) - 1), robHelper(nums, 0, len(nums) - 2))
}

func robHelper(nums []int, start int, end int) int {
    if end == start {
        return nums[start]
    }
    
    len := end - start + 1
    prev, cur := 0, nums[start]

    for i := 1; i < len; i++ {
        next := max(cur, prev + nums[start + i])
        prev = cur
        cur = next
    }

    return cur
}