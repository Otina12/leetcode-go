func maxSumAfterPartitioning(arr []int, k int) int {
    memo := make([]int, len(arr))
    var solve func(int) int

    solve = func(i int) int {
        if i == len(arr) {
            return 0
        }

        if memo[i] != 0 {
            return memo[i]
        }

        maxNum := 0
        ans := 0

        for j := i; j < len(arr) && j < i + k; j++ {
            maxNum = max(maxNum, arr[j])
            curCnt := j - i + 1

            ans = max(ans, maxNum * curCnt + solve(j + 1))
        }

        memo[i] = ans
        return ans
    }

    return solve(0)
}