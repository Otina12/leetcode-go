func numTrees(n int) int {
    var helper func(int, int) int

    helper = func(start, end int) int {
        if start >= end {
            return 1
        }

        total := 0
        for i := start; i <= end; i++ {
            left := helper(start, i - 1)
            right := helper(i + 1, end)
            total += left * right
        }

        return total
    }
    
    return helper(1, n)
}