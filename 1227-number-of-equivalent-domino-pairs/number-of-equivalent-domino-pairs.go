func numEquivDominoPairs(dominoes [][]int) int {
    freqMap := make(map[int]int)

    for _, domino := range dominoes {
        key := 10 * min(domino[0], domino[1]) + max(domino[0], domino[1])
        freqMap[key] += 1
    }

    res := 0

    for _, frequency := range freqMap {
        res += frequency * (frequency - 1) / 2
    }
    
    return res
}