func distributeCookies(cookies []int, k int) int {
    distribution := make([]int, k)
    var backtrack func(int)
    res := 8 * 100000

    backtrack = func(cookieIdx int) {
        if cookieIdx == len(cookies) {
            maxUnfairness := 0
            for _, val := range distribution {
                maxUnfairness = max(maxUnfairness, val)
            }
            res = min(res, maxUnfairness)
            return
        }

        for i := 0; i < k; i++ {
            distribution[i] += cookies[cookieIdx]

            if distribution[i] < res {
                backtrack(cookieIdx + 1)
            }

            distribution[i] -= cookies[cookieIdx]
        }
    }

    backtrack(0)
    return res
}