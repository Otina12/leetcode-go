type memoKey struct {
    num1 int
    num2 int
    num3 int
}

func dieSimulator(n int, rollMax []int) int {
    MOD := 1_000_000_007
    memo := make(map[memoKey]int)
    var dfs func(int, int, int) int

    dfs = func(rolls int, num int, numLeft int) int {
        if rolls == n {
            return 1
        }

        key := memoKey{rolls, num, numLeft}
        if val, exists := memo[key]; exists {
            return val
        }

        res := 0

        if numLeft > 0 {
            res += dfs(rolls + 1, num, numLeft - 1)
        }

        for i := 0; i < 6; i++ {
            if i != num {
                res += dfs(rolls + 1, i, rollMax[i] - 1)
            }
        }

        res %= MOD
        memo[key] = res
        return res
    }

    return dfs(0, -1, 0)
}