func getWordsInLongestSubsequence(words []string, groups []int) []string {
    n := len(words)
    dp := make([][]int, n)
    
    for i := 0; i < n; i++ {
        dp[i] = []int{i}
    }

    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            if groups[i] == groups[j] || len(words[i]) != len(words[j]) {
                continue
            }

            if len(dp[j]) >= len(dp[i]) && hammingDistance(words[i], words[j]) == 1  {
                newSubseq := make([]int, len(dp[j]))
                copy(newSubseq, dp[j])
                newSubseq = append(newSubseq, i)
                dp[i] = newSubseq
            }
        }
    }

    maxLen, maxLenIdx := -1, -1

    for i := 0; i < n; i++ {
        if len(dp[i]) > maxLen {
            maxLen = len(dp[i])
            maxLenIdx = i
        }
    }

    var res []string

    for _, wordIdx := range dp[maxLenIdx] {
        res = append(res, words[wordIdx])
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