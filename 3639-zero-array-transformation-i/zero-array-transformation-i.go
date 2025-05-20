func isZeroArray(nums []int, queries [][]int) bool {
    n := len(nums)
    differenceArray := make([]int, n)

    for _, query := range queries {
        from, to := query[0], query[1]
        differenceArray[from] += 1
        if to < n - 1 {
            differenceArray[to + 1] -= 1
        }
    }

    prefixSum := 0
    for i, num := range nums {
        prefixSum += differenceArray[i]
        if prefixSum < num {
            return false
        }
    }

    return true
}