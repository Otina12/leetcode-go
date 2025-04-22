func numberOfSubmatrices(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    prefixCount := make([][][]int, m + 1)
    res := 0

    for i := 0; i <= m; i++ {
        prefixCount[i] = make([][]int, n+1)
        for j := 0; j <= n; j++ {
            prefixCount[i][j] = make([]int, 2)
        }
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if grid[i-1][j-1] == 'X' {
                prefixCount[i][j] = []int {1, 0}
            } else if grid[i-1][j-1] == 'Y' {
                prefixCount[i][j] = []int {0, 1}
            }
            
            prefixCount[i][j][0] += prefixCount[i-1][j][0] + prefixCount[i][j-1][0] - prefixCount[i-1][j-1][0]
            prefixCount[i][j][1] += prefixCount[i-1][j][1] + prefixCount[i][j-1][1] - prefixCount[i-1][j-1][1]

            if prefixCount[i][j][0] > 0 && prefixCount[i][j][0] == prefixCount[i][j][1] {
                res += 1
            }
        }
    }

    return res
}