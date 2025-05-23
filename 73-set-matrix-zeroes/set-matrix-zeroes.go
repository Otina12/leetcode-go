func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    zeroInFirstRow, zeroInFirstColumn := false, false

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0

                if i == 0 {
                    zeroInFirstRow = true
                }
                if j == 0 {
                    zeroInFirstColumn = true
                }
            }
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if zeroInFirstRow {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }

    if zeroInFirstColumn {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}