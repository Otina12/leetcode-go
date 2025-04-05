func subsetXORSum(nums []int) int {
    var res int = 0

    for _, num := range nums {
        res |= num
    }

    return res << (len(nums) - 1)
}