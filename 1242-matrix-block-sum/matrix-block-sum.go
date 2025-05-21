func matrixBlockSum(mat [][]int, k int) [][]int {
    m, n := len(mat), len(mat[0])
    prefixSumMatrix := make([][]int, m + 1)

    for i := 0; i <= m; i++ {
        prefixSumMatrix[i] = make([]int, n + 1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            left := prefixSumMatrix[i][j-1]
            upper := prefixSumMatrix[i-1][j]
            diag := prefixSumMatrix[i-1][j-1]
            prefixSumMatrix[i][j] = mat[i-1][j-1] + left + upper - diag
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            r1 := max(0, i - k)
            c1 := max(0, j - k)
            r2 := min(m, i + k + 1)
            c2 := min(n, j + k + 1)

            mat[i][j] = prefixSumMatrix[r2][c2] -
                        prefixSumMatrix[r1][c2] -
                        prefixSumMatrix[r2][c1] +
                        prefixSumMatrix[r1][c1]
        }
    }

    return mat
}