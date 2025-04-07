func canPartition(nums []int) bool {
    var sum int = 0

	for _, num := range nums {
		sum += num
	}

	if sum % 2 != 0 {
		return false
	}

	var target int = sum / 2
	var dp []bool = make([]bool, target + 1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}