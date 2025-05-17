func minNumberOperations(target []int) int {
    res := 0
    curSubarrayMax := target[0]

    for i := 1; i < len(target); i++ {
        if target[i] < target[i-1] {
            res += curSubarrayMax - target[i]
            curSubarrayMax = target[i]
        } else {
            curSubarrayMax = max(curSubarrayMax, target[i])
        }
    }

    res += curSubarrayMax
    return res
}