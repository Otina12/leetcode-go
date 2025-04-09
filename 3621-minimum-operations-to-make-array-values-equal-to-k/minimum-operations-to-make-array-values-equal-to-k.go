func minOperations(nums []int, k int) int {
    var uniqueNums = make(map[int]bool)

    for _, num := range nums {
        uniqueNums[num] = true
    }

    var res = 0

    for num, _ := range uniqueNums {
        if num < k {
            return -1
        } else if num > k {
            res += 1
        }
    }

    return res
}