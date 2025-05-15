func getLongestSubsequence(words []string, groups []int) []string {
    lastDigit := groups[0]
    res := []string{words[0]}

    for i := 1; i < len(groups); i++ {
        if abs(groups[i] - lastDigit) == 1 {
            lastDigit = groups[i]
            res = append(res, words[i])
        }
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}