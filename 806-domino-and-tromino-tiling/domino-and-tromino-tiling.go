type memoPair struct {
    num1 int
    num2 bool
}

func numTilings(n int) int {
    var MOD = 1000000007
    var memo = make(map[memoPair]int)
    var helper func(int, bool) int

    helper = func(column int, oddLeft bool) int {
        if column > n + 1 {
            return 0
        }
        if column == n + 1 {
            if !oddLeft {
                return 1
            }
            return 0
        }

        pair := memoPair{column, oddLeft}
        if val, exists := memo[pair]; exists {
            return val
        }

        res := 0
        if oddLeft {
            res += helper(column + 1, true) // place domino
            res += helper(column + 2, false) // place tromino
        } else {
            res += helper(column + 1, false) // place domino
            res += helper(column + 2, false) // place 2 dominoes
            res += 2 * helper(column + 1, true) // place tromino
        }

        res %= MOD
        memo[pair] = res
        return res
    }

    return helper(1, false)
}