func takeCharacters(s string, k int) int {
    letterCount := []int {0, 0, 0} // 0: aCount, 1: bCount, 2: cCount

    for _, c := range s {
        letterCount[c - 'a'] += 1
    }

    if letterCount[0] < k || letterCount[1] < k || letterCount[2] < k {
        return -1
    }

    curCount := []int {0, 0, 0}
    left := 0
    maxLeftoverCnt := 0

    for right, c := range s {
        charIndex := c - 'a'
        curCount[charIndex] += 1

        for left <= right && curCount[charIndex] > letterCount[charIndex] - k {
            curCount[s[left] - 'a'] -= 1
            left += 1
        }

        maxLeftoverCnt = max(maxLeftoverCnt, right - left + 1)
    }

    return len(s) - maxLeftoverCnt
}