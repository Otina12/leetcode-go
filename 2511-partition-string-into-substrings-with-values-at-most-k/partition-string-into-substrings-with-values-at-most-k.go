func minimumPartition(s string, k int) int {
    if k < 10 {
        for _, digit := range s {
            if int(digit - '0') > k {
                return -1
            }
        }
    }
    
    var partition func(int, int) int

    partition = func(i int, prevNum int) int {
        if i == len(s) {
            return 0
        }

        curNum := 10 * prevNum + int(s[i] - '0') // always less than k

        if i < len(s) - 1 && curNum * 10 + int(s[i+1] - '0') <= k {
            return partition(i + 1, curNum)
        }

        return 1 + partition(i + 1, 0)
    }

    return partition(0, 0)
}