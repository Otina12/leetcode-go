func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    n := len(capacity)
    emptySpaces := make([]int, n)

    for i := 0; i < n; i++ {
        emptySpaces[i] = capacity[i] - rocks[i]
    }

    sort.Ints(emptySpaces)
    bagIdx := 0
    for emptySpaces[bagIdx] == 0 && bagIdx < n {
        bagIdx += 1
    }
    
    for additionalRocks > 0 && bagIdx < n {
        if additionalRocks < emptySpaces[bagIdx] {
            break
        }
        
        additionalRocks -= emptySpaces[bagIdx]
        bagIdx += 1
    }

    return bagIdx
}