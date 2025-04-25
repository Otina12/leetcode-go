func findBall(grid [][]int) []int {
    ballsCnt := len(grid[0])
    res := make([]int, ballsCnt)

    for col := 0; col < ballsCnt; col++ {
        res[col] = getBallFallColumn(grid, col)
    }

    return res
}

func getBallFallColumn(grid [][]int, ballColumn int) int {
    m, n := len(grid), len(grid[0])
    row, col := 0, ballColumn

    for row < m {
        newCol := col + grid[row][col]

        if newCol < 0 || newCol >= n || grid[row][col] != grid[row][newCol] {
            col = -1
            break
        }

        col = newCol
        row += 1
    }

    return col
}