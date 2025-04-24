func numPairsDivisibleBy60(time []int) int {
    remainderFreq := make(map[int]int)
    res := 0

    for _, duration := range time {
        remainderFreq[duration % 60] += 1
    }

    for durationRemainder, freq := range remainderFreq {
        if durationRemainder == 30 || durationRemainder == 0 {
            res += freq * (freq - 1)
        } else {
            res += freq * remainderFreq[60 - durationRemainder]
        }

        
    }

    return res / 2
}