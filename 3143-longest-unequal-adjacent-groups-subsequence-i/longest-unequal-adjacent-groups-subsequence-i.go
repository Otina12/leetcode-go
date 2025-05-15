type pair struct {
    indices []int
    length int
}

func getLongestSubsequence(words []string, groups []int) []string {
    maxEndingAt := [2]pair{}

    for i := 0; i < len(groups); i++ {
        maxEndingAt[groups[i]].indices = append(maxEndingAt[1-groups[i]].indices, i)
        maxEndingAt[groups[i]].length = maxEndingAt[1-groups[i]].length + len(words[i])
    }

    maxI := 0
    if maxEndingAt[1].length > maxEndingAt[0].length {
        maxI = 1
    }

    res := []string{}

    for _, index := range maxEndingAt[maxI].indices {
        res = append(res, words[index])
    }

    return res
}