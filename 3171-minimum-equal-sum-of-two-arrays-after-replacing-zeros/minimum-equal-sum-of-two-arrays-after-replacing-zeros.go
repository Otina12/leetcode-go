func minSum(nums1 []int, nums2 []int) int64 {
    sum1, sum2 := 0, 0
    zeroCnt1, zeroCnt2 := 0, 0

    for _, num := range nums1 {
        if num == 0 {
            zeroCnt1 += 1
        }
        sum1 += num
    }

    for _, num := range nums2 {
        if num == 0 {
            zeroCnt2 += 1
        }
        sum2 += num
    }

    if sum1 == sum2 && zeroCnt1 * zeroCnt2 == 0 && (zeroCnt1 != 0 || zeroCnt2 != 0) {
        return -1
    }

    if sum1 < sum2 && (zeroCnt1 == 0 || zeroCnt2 == 0 && sum2 - sum1 < zeroCnt1) {
        return -1
    }

    if sum2 < sum1 && (zeroCnt2 == 0 || zeroCnt1 == 0 && sum1 - sum2 < zeroCnt2) {
        return -1
    }

    return int64(max(sum1 + zeroCnt1, sum2 + zeroCnt2))
}