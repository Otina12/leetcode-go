func makesquare(matchsticks []int) bool {
    perimeter := 0

    for _, matchstick := range matchsticks {
        perimeter += matchstick
    }

    if perimeter % 4 != 0 {
        return false
    }

    sideLength := perimeter / 4
    sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
    used := make([]bool, len(matchsticks))
    var dfs func(int, int, int) bool

    dfs = func(start int, curSum int, sides int) bool {
        if sides == 3 {
            return true
        }

        if curSum == sideLength {
            return dfs(0, 0, sides + 1)
        }

        for i := start; i < len(matchsticks); i++ {
            if used[i] || curSum + matchsticks[i] > sideLength {
                continue
            }

            used[i] = true
            if dfs(i + 1, curSum + matchsticks[i], sides) {
                return true
            }
            used[i] = false

            if curSum == 0 {
                break
            }
        }

        return false
    }

    return dfs(0, 0, 0)
}