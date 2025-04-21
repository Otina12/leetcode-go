func numberOfArrays(differences []int, lower int, upper int) int {
    curStart, curEnd := lower, upper

    for _, diff := range differences {
        curStart += diff
        curEnd += diff

        curStart = max(curStart, lower)
        curEnd = min(curEnd, upper)
    }

    return max(0, curEnd - curStart + 1)
}