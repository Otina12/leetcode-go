func countCompleteDayPairs(hours []int) int64 {
    remainderFreq := make(map[int]int64)
    var res int64 = 0

    for _, duration := range hours {
        remainderFreq[duration % 24] += 1
    }

    for durationRemainder, freq := range remainderFreq {
        if durationRemainder == 0 || durationRemainder == 12 {
            res += freq * (freq - 1)
        } else {
            res += freq * remainderFreq[24 - durationRemainder]
        }

        
    }

    return res / 2
}