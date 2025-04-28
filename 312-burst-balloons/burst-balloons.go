type Pair struct {
    num1 int
    num2 int
}

func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

    memo := make(map[Pair]int)
    var dfs func(int, int) int

    dfs = func(left int, right int) int {
        if left > right {
            return 0
        }

        pair := Pair{left, right}

        if val, exists := memo[pair]; exists {
            return val
        }

        memo[pair] = 0

        for i := left; i <= right; i++ {
            memo[pair] = max(memo[pair], nums[left-1] * nums[i] * nums[right+1] + dfs(left, i-1) + dfs(i+1, right))
        }

        return memo[pair]
    }

    return dfs(1, len(nums) - 2)
}