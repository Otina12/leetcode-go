func restoreArray(adjacentPairs [][]int) []int {
    n := len(adjacentPairs)
    pairs := make(map[int][]int)

    for _, pair := range adjacentPairs {
        a, b := pair[0], pair[1]
        pairs[a] = append(pairs[a], b)
        pairs[b] = append(pairs[b], a)
    }

    start := 0

    for num, neighbors := range pairs {
        if len(neighbors) == 1 {
            start = num
            break
        }
    }

    res := make([]int, 0, n + 1)
    prev, cur := -1, start

    for len(res) < n + 1 {
        res = append(res, cur)

        for _, next := range pairs[cur] {
            if next != prev {
                prev = cur
                cur = next
                break
            }
        }
    }

    return res
}