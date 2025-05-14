func nextBeautifulNumber(n int) int {
    var backtrack func(int, map[int]int)
    var isBalanced func(int) bool
    var res int = 1666666

    isBalanced = func(num int) bool {
        freq := [7]int{}
        
        for num > 0 {
            freq[num % 10] += 1
            num /= 10
        }

        for i := 0; i < len(freq); i++ {
            if freq[i] > 0 && i != freq[i] {
                return false
            }
        }
        
        return true
    }

    backtrack = func(currentNum int, counter map[int]int) {
        if currentNum > n && isBalanced(currentNum) {
            res = min(res, currentNum)
            return
        }

        if currentNum > res {
            return
        }

        for d := 1; d < 7; d++ {
            if counter[d] == d {
                continue
            }

            counter[d] += 1
            backtrack(10 * currentNum + d, counter)
            counter[d] -= 1
        }
    }

    backtrack(0, make(map[int]int))
    return res
}