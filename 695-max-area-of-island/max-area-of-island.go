func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var maxArea int
    var visited = make([][]bool, m)

    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }
    
    var dfs func(int, int) int
    dfs = func(i int, j int) int {
        if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == 0 {
            return 0
        }

        visited[i][j] = true
        return 1 + dfs(i + 1, j) + dfs(i, j + 1) + dfs(i - 1, j) + dfs(i, j - 1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                maxArea = max(maxArea, dfs(i, j))
            }
        }
    }

    return maxArea
}