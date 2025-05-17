func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    changes := make([]int, n)

    for i := 0; i < n; i++ {
        changes[i] = gas[i] - cost[i]
    }

    i, j := 0, n - 1
    totalGas := 0

    for i <= j {
        totalGas += changes[i]
        for i < j && totalGas < 0 {
            totalGas += changes[j]
            j -= 1
        }
        i += 1
    }

    if totalGas < 0 {
        return -1
    }

    return (j + 1) % n
}