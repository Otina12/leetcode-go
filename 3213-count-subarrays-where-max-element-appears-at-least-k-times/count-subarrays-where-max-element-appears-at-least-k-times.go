func countSubarrays(nums []int, k int) int64 {
    maxNum := getMax(nums)
    left, maxNumCount := 0, 0
    res := int64(0)

    for right := 0; right < len(nums); right++ {
        if nums[right] == maxNum {
            maxNumCount += 1

            if maxNumCount >= k {
                for left < right && nums[left] != maxNum {
                    left += 1
                }
                
                left += 1
            }
        }

        if maxNumCount >= k {
            res += int64(left)
        }
    }

    return res
}

func getMax(nums []int) int {
    maxNum := 0

    for _, num := range nums {
        maxNum = max(maxNum, num)
    }

    return maxNum
}