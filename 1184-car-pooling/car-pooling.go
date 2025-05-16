func carPooling(trips [][]int, capacity int) bool {
    changes := make([]int, 1001)

    for _, trip := range trips {
        passengerCnt, from, to := trip[0], trip[1], trip[2]
        changes[from] += passengerCnt
        changes[to] -= passengerCnt
    }

    curPassengers := 0

    for _, change := range changes {
        curPassengers += change
        if curPassengers > capacity {
            return false
        }
    }

    return true
}