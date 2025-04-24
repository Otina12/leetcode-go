func maxDistance(path string, k int) int {
	res := 0
    totalDirs := 0
    dirFreq := make(map[rune]int)

    for _, dir := range path {
        dirFreq[dir] += 1
        totalDirs += 1

        opposingDirections := min(dirFreq['N'], dirFreq['S']) + min(dirFreq['E'], dirFreq['W'])
        currentDistance := totalDirs - 2 * opposingDirections

        res = max(res, currentDistance + 2 * min(k, opposingDirections))
    }
    
    return res
}