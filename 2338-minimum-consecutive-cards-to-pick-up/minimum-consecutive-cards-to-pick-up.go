func minimumCardPickup(cards []int) int {
    lastOccurence := make(map[int]int)
    res := 100001

    for i, num := range cards {
        if idx, exists := lastOccurence[num]; exists {
            res = min(res, i - idx + 1)
        }

        lastOccurence[num] = i
    }

    if res == 100001 {
        return -1
    }

    return res
}