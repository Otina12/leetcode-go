func minDominoRotations(tops []int, bottoms []int) int {
    n := len(tops)
    minRotations := 20001
    values := []int {tops[0], bottoms[0]}

    for _, value := range values {
        canMatchValues := true
        topRotations, bottomRotations := 0, 0

        for i := 0; i < n; i++ {
            if tops[i] != value && bottoms[i] != value {
                canMatchValues = false
                break
            }

            if tops[i] != value {
                topRotations += 1
            }
            if bottoms[i] != value {
                bottomRotations += 1
            }
        }

        if canMatchValues {
            minRotations = min(minRotations, min(topRotations, bottomRotations))
        }
    }

    if minRotations == 20001 {
        return -1
    }

    return minRotations
}