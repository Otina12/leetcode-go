func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)
    numNeeded := 1

    for i := 0; i < len(arr); i++ {
        if arr[i] >= numNeeded {
            numNeeded += 1
        }
    }

    return numNeeded - 1
}