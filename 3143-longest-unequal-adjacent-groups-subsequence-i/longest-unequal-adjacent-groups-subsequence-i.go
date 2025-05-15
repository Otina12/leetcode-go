func getLongestSubsequence(words []string, groups []int) []string {
    maxEndingAt := make([][]int, 2)
    maxEndingAt[0] = []int{}
    maxEndingAt[1] = []int{}

    for i := 0; i < len(groups); i++ {
        maxEndingAt[groups[i]] = append(maxEndingAt[1-groups[i]], i)
    }

    var res []string
    var maxI int = 0
    
    if len(maxEndingAt[1]) > len(maxEndingAt[0]) {
        maxI = 1
    }

    for _, wordI := range maxEndingAt[maxI] {
        res = append(res, words[wordI])
    }

    return res
}