func maxSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    prefixSumMatrix := make([][]int, m + 1)

    for i := 0; i <= m; i++ {
        prefixSumMatrix[i] = make([]int, n + 1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            left := prefixSumMatrix[i][j-1]
            upper := prefixSumMatrix[i-1][j]
            diag := prefixSumMatrix[i-1][j-1]
            prefixSumMatrix[i][j] = grid[i-1][j-1] + left + upper - diag
        }
    }

    res := 0

    for i := 2; i < m; i++ {
        for j := 2; j < n; j++ {
            r1, c1 := i - 2, j - 2
            r2, c2 := i + 1, j + 1

            blockSum := prefixSumMatrix[r2][c2] -
                        prefixSumMatrix[r1][c2] -
                        prefixSumMatrix[r2][c1] +
                        prefixSumMatrix[r1][c1]

            res = max(res, blockSum - grid[i-1][j] - grid[i-1][j-2])
        }
    }

    return res
}