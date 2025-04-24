func countSubmatrices(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    prefixSum := make([][]int, m + 1)

    for i := 0; i <= m; i++ {
        prefixSum[i] = make([]int, n + 1)
    }

    res := 0

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            prefixSum[i][j] = grid[i-1][j-1] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
            if prefixSum[i][j] <= k {
                res += 1
            }
        }
    }

    return res
}