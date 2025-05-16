func maximumLength(nums []int, k int) int {
    n := len(nums)
    dp := make([][][]int, n + 1)

    for i := 0; i <= n; i++ {
        dp[i] = make([][]int, n + 1)
        for j := 0; j <= n; j++ {
            dp[i][j] = make([]int, k + 1)
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= n; j++ {
            for curK := 0; curK <= k; curK++ {
                skip := dp[i+1][j][curK]
                take := -500

                if j == n {
                    take = 1 + dp[i+1][i][curK]
                } else {
                    if nums[i] == nums[j] {
                        take = 1 + dp[i+1][i][curK]
                    } else if curK + 1 <= k {
                        take = 1 + dp[i+1][i][curK+1]
                    }
                }

                dp[i][j][curK] = max(skip, take)
            }
        }
    }

    return dp[0][n][0]
}