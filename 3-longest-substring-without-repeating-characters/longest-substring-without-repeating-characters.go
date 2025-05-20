func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int)
    left := 0
    res := 0

    for right := 0; right < len(s); right++ {
        if lastOccur, exists := seen[s[right]]; exists {
            left = max(left, lastOccur + 1)
        }

        seen[s[right]] = right
        res = max(res, right - left + 1)
    }

    return res
}