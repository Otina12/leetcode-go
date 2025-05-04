func minWindow(s string, t string) string {
    frequencyMap, patternMap := make(map[rune]int), make(map[rune]int)

    for _, ch := range t {
        patternMap[ch] += 1
    }

    left, maxLen := 0, len(s) + 1
    res := ""
    windowIsSubstring := false

    for right, ch := range s {
        frequencyMap[ch] += 1

        if !windowIsSubstring {
            windowIsSubstring = isSubstring(frequencyMap, patternMap)
        }

        if windowIsSubstring {
            for {
                curRune := rune(s[left])
                freq, exists := patternMap[curRune]
                if !exists || frequencyMap[curRune] > freq {
                    frequencyMap[curRune] -= 1
                    left += 1
                } else {
                    break
                }
            }

            if right - left + 1 < maxLen {
                maxLen = right - left + 1
                res = s[left:right+1]
            }
        }
    }

    return res
}

func isSubstring(strFreq map[rune]int, patternFreq map[rune]int) bool {
    for key, patternVal := range patternFreq {
        if strVal, exists := strFreq[key]; !exists || strVal < patternVal {
            return false
        }
    }

    return true
}