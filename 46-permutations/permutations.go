func permute(nums []int) [][]int {
    var allPermutations [][]int
    var generatePermutations func(int)

    generatePermutations = func(i int) {
        if i == len(nums) {
            permutation := make([]int, len(nums))
            copy(permutation, nums)
            allPermutations = append(allPermutations, permutation)
            return
        }

        for j := i; j < len(nums); j++ {
            nums[i], nums[j] = nums[j], nums[i]
            generatePermutations(i + 1)
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    generatePermutations(0)
    return allPermutations
}