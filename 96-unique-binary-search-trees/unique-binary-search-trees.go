type pair struct {
    num1, num2 int
}

func numTrees(n int) int {
    memo := make(map[pair]int)
    var helper func(int, int) int

    helper = func(start, end int) int {
        if start >= end {
            return 1
        }

        key := pair{start, end}
        if val, exists := memo[key]; exists {
            return val
        }

        total := 0
        for i := start; i <= end; i++ {
            left := helper(start, i - 1)
            right := helper(i + 1, end)
            total += left * right
        }

        memo[key] = total
        return total
    }
    
    return helper(1, n)
}