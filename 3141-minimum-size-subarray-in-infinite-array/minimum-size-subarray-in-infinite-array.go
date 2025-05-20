func minSizeSubarray(nums []int, target int) int {
    n := len(nums)
    nums2 := make([]int, 2*n)
    totalSum := 0

    for _, num := range nums {
        totalSum += num
    }

    wholeArrCnt := target / totalSum
    target %= totalSum

    copy(nums2, nums)

    for i := 0; i < n; i++ {
        nums2[n+i] = nums[i]
    }

    curWindowSum := 0
    left, res := 0, n

    for right, num := range nums2 {
        curWindowSum += num

        for left <= right && curWindowSum > target {
            curWindowSum -= nums2[left]
            left += 1
        }

        if curWindowSum == target {
            res = min(res, right - left + 1)
        }
    }

    if res == n {
        return -1
    }

    return res + n * wholeArrCnt
}