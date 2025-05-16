func getWordsInLongestSubsequence(words []string, groups []int) []string {
    n := len(words)
    dp := make([]int, n)
    prev := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = 1
        prev[i] = -1

        for j := 0; j < i; j++ {
            if groups[i] == groups[j] || len(words[i]) != len(words[j]) {
                continue
            }

            if dp[j] >= dp[i] && hammingDistance(words[i], words[j]) == 1  {
                dp[i] = dp[j] + 1
                prev[i] = j
            }
        }
    }

    maxLen, lastIdx := 0, -1

    for i := 0; i < n; i++ {
        if dp[i] > maxLen {
            maxLen = dp[i]
            lastIdx = i
        }
    }

    var indices []int
    for lastIdx != -1 {
        indices = append(indices, lastIdx)
        lastIdx = prev[lastIdx]
    }

    var res []string
    for i := len(indices) - 1; i >= 0; i-- {
        res = append(res, words[indices[i]])
    }

    return res
}

func hammingDistance(word1 string, word2 string) int { // Invariant: word1 and word2 lengths are equal.
    distance := 0

    for i := 0; i < len(word1); i++ {
        if word1[i] != word2[i] {
            distance += 1
        }
    }

    return distance
}