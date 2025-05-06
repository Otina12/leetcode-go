func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        res += (longestPalindromeLength(s, i, i) + 1) / 2 // odd length

        if i < len(s) - 1 && s[i] == s[i+1] {
            res += longestPalindromeLength(s, i, i + 1) / 2 // even length
        }
    }

    return res
}

func longestPalindromeLength(str string, start1 int, start2 int) int {
    for start1 >= 0 && start2 < len(str) && str[start1] == str[start2] {
        start1 -= 1
        start2 += 1
    }

    start1 += 1
    start2 -= 1

    return start2 - start1 + 1
}