var _res int
var _nums []int

func subsetXORSum(nums []int) int {
    _res = 0
    _nums = nums

    helper(0, 0)
    return _res
}

func helper(i int, cur_xor int) {
    if i == len(_nums) {
        _res += cur_xor
        return
    }

    helper(i + 1, cur_xor)
    helper(i + 1, cur_xor ^ _nums[i])
}