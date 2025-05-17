func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    i, j := 0, n - 1
    totalGas := 0

    for i <= j {
        totalGas += gas[i] - cost[i]
        for i < j && totalGas < 0 {
            totalGas += gas[j] - cost[j]
            j -= 1
        }
        i += 1
    }

    if totalGas < 0 {
        return -1
    }

    return (j + 1) % n
}