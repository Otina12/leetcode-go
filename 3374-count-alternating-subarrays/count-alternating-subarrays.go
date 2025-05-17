func countAlternatingSubarrays(nums []int) int64 {
    res := 1
    left := 0

    for right := 1; right < len(nums); right++ {
        if nums[right - 1] == nums[right] {
            left = right
        }

        res += right - left + 1
    }

    return int64(res)
}