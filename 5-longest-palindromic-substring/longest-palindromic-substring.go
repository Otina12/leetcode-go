func longestPalindrome(s string) string {
    maxLen := 0
    res := ""

    for i := 0; i < len(s); i++ {
        // odd length
        l, r := detectLongestPalindrome(s, i, i)

        if r - l + 1 > maxLen {
            maxLen = r - l + 1
            res = s[l:r+1]
        }

        // even length
        l, r = detectLongestPalindrome(s, i, i + 1)

        if r - l + 1 > maxLen {
            maxLen = r - l + 1
            res = s[l:r+1]
        }
    }

    return res
}

func detectLongestPalindrome(str string, start1 int, start2 int) (int, int) {
    for start1 >= 0 && start2 < len(str) && str[start1] == str[start2] {
        start1 -= 1
        start2 += 1
    }

    return start1 + 1, start2 - 1
}