func partitionString(s string) int {
    visitedSet := make(map[rune]struct{})
    res := 0

    for _, ch := range s {
        if _, exists := visitedSet[ch]; exists {
            visitedSet = make(map[rune]struct{})
            res += 1
        }

        visitedSet[ch] = struct{}{}
    }

    return res + 1
}