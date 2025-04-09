func wordBreak(s string, wordDict []string) bool {
    var wordSet map[string]bool = make(map[string]bool)
    var memo []int = make([]int, len(s))

    for _, word := range wordDict {
        wordSet[word] = true
    }

    var canBreak func(int) bool

    canBreak = func(startI int) bool {
        if startI == len(s) {
            return true
        }

        if memo[startI] == 0 {
            memo[startI] = -1

            for endI := startI; endI < len(s); endI++ {
                word := s[startI:endI+1]
                _, isInDictionary := wordSet[word]

                if isInDictionary && canBreak(endI + 1) {
                    memo[startI] = 1
                    break
                }
            }
        }

        return memo[startI] == 1
    }
    
    return canBreak(0)
}