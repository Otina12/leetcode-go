func numOfSubarrays(arr []int, k int, threshold int) int {
    left, curSum := 0, 0
    res := 0

    for right, num := range arr {
        curSum += num

        if right - left + 1 > k {
            curSum -= arr[left]
            left += 1
        }

        if right - left + 1 == k {
            average := float64(curSum) / float64(right - left + 1)
            if int(average) >= threshold {
                res += 1
            }
        }
    }

    return res
}