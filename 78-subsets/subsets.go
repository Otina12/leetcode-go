func subsets(nums []int) [][]int {
    var allSubsets [][]int
    var currentSubset []int
    var generateSubsets func(int)

    generateSubsets = func(i int) {
        if i == len(nums) {
            subset := make([]int, len(currentSubset))
            copy(subset, currentSubset)
            allSubsets = append(allSubsets, subset)
            return
        }

        generateSubsets(i + 1)
        currentSubset = append(currentSubset, nums[i])
        generateSubsets(i + 1)
        currentSubset = currentSubset[:len(currentSubset) - 1]
    }

    generateSubsets(0)
    return allSubsets
}