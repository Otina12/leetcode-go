func maxUniqueSplit(s string) int {
    stringSet := make(map[string]bool)
    var backtrack func(int) int

    backtrack = func(i int) int {
        if i == len(s) {
            return 0
        }

        ans := 0

        for right := i + 1; right <= len(s); right++ {
            substring := s[i:right]
            if !stringSet[substring] {
                stringSet[substring] = true
                ans = max(ans, 1 + backtrack(right))
                delete(stringSet, substring)
            }
        }

        return ans
    }

    return backtrack(0)
}