func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    prefixCountFreq := make(map[int]int64)
    prefixCountFreq[0] = 1
    prefixCount := 0
    var res int64 = 0

    for _, num := range nums {
        if num % modulo == k {
            prefixCount += 1
        }

        prefixCount %= modulo

        if _, exists := prefixCountFreq[(prefixCount - k + modulo) % modulo]; exists {
            res += prefixCountFreq[(prefixCount - k + modulo) % modulo]
        }

        prefixCountFreq[prefixCount] += 1
    }

    return res
}