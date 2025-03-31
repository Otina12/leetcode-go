import (
    "slices"
)

func putMarbles(weights []int, k int) int64 {
    var splits []int = []int {}

    for i := 1; i < len(weights); i++ {
        splits = append(splits, weights[i-1] + weights[i])
    }

    slices.Sort(splits)
    var minScore, maxScore int64 = 0, 0

    for i := 0; i < k - 1; i++ {
        minScore += int64(splits[i])
        maxScore += int64(splits[len(splits)-1-i])
    }

    return maxScore - minScore
}