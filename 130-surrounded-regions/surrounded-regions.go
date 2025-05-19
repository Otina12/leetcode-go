func solve(board [][]byte)  {
    m, n := len(board), len(board[0])
    var markSafeRegion func(int, int)

    markSafeRegion = func(row int, col int) {
        if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != 'O' {
            return
        }

        board[row][col] = '#'

        markSafeRegion(row + 1, col)
        markSafeRegion(row - 1, col)
        markSafeRegion(row, col + 1)
        markSafeRegion(row, col - 1)
    }

    for i := 0; i < m; i++ {
        markSafeRegion(i, 0)
        markSafeRegion(i, n - 1)
    }

    for j := 0; j < n; j++ {
        markSafeRegion(0, j)
        markSafeRegion(m - 1, j)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == '#' {
                board[i][j] = 'O'
            }
        }
    }
}