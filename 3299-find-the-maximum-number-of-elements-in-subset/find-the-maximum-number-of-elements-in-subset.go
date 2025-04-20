func maximumLength(nums []int) int {
    freq := make(map[int]int)
    res := 0

    for _, num := range nums {
        freq[num] += 1
    }

    if freq[1] & 1 == 0 {
        res = max(res, freq[1] - 1)
    } else {
        res = max(res, freq[1])
    }

    for key, _ := range freq {
        if key == 1 {
            continue
        }
        
        num := key
        curSize := 0

        for {
            cnt, exists := freq[num]
            if !exists {
                res = max(res, curSize - 1) // count last seen element as x^k
                break
            }

            if cnt >= 2 {
                curSize += 2
                num = num * num
            } else {
                res = max(res, curSize + 1)
                break
            }
        }
    }

    return res
}