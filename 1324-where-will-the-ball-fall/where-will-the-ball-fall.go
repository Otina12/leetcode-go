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
    r, c := 0, ballColumn

    for r < m {
        if grid[r][c] == 1 {
            if c == n - 1 || grid[r][c+1] == -1 {
                return -1
            }
            c += 1
        } else if grid[r][c] == -1 {
            if c == 0 || grid[r][c-1] == 1 {
                return -1
            }
            c -= 1
        }

        r += 1
    }

    return c
}