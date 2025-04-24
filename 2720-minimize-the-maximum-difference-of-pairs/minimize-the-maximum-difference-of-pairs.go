func minimizeMax(nums []int, p int) int {
    sort.Ints(nums)
    left := 0
    right := nums[len(nums) - 1] - nums[0]
    res := 0

    for left <= right {
        mid := left + (right - left) / 2

        if canFormKPairs(nums, mid, p) {
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return res
}

func canFormKPairs(nums []int, maxSum int, k int) bool {
    i := 1
    cnt := 0

    for i < len(nums) {
        if nums[i] - nums[i-1] <= maxSum {
            i += 2
            cnt += 1
        } else {
            i += 1
        }

        if cnt == k {
            return true
        }
    }

    return false
}