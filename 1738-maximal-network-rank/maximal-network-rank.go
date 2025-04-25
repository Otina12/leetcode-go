func maximalNetworkRank(n int, roads [][]int) int {
    adjList := make([]map[int]bool, n)

    for i := 0; i < n; i++ {
        adjList[i] = make(map[int]bool)
    }

    for _, road := range roads {
        a, b := road[0], road[1]

        adjList[a][b] = true
        adjList[b][a] = true
    }

    res := 0
    
    for a := 0; a < n; a++ {
        for b := a + 1; b < n; b++ {
            if _, exists := adjList[a][b]; exists {
                res = max(res, len(adjList[a]) + len(adjList[b]) - 1)
            } else {
                res = max(res, len(adjList[a]) + len(adjList[b]))
            }
        }
    }

    return res
}