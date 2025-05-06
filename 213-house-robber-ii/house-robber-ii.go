func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    lastHouse := nums[len(nums) - 1]
    
    nums = nums[:len(nums) - 1]
    withoutLastHouse := robHelper(nums)

    nums = append(nums, lastHouse)
    nums = nums[1:]

    withoutFirstHouse := robHelper(nums)

    return max(withoutLastHouse, withoutFirstHouse)
}

func robHelper(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    prev, cur := 0, nums[0]

    for i := 1; i < len(nums); i++ {
        next := max(cur, prev + nums[i])
        prev = cur
        cur = next
    }

    return cur
}