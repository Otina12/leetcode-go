func matrixSum(nums [][]int) int {
    m, n := len(nums), len(nums[0])
    score := 0

    for i := 0; i < m; i++ {
        sort.Ints(nums[i])
    }

    for j := 0; j < n; j++ {
        maxInColumn := 0
        for i := 0; i < m; i++ {
            maxInColumn = max(maxInColumn, nums[i][j])
        }
        score += maxInColumn
    }

    return score
}