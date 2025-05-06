func buildArray(nums []int) []int {
    base := 1024

    for i := 0; i < len(nums); i++ {
        nums[i] += base * (nums[nums[i]] % base)
    }

    for i := 0; i < len(nums); i++ {
        nums[i] /= base
    }

    return nums
}