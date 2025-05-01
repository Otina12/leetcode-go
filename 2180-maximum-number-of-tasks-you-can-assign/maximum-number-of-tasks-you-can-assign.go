func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
    var canCompleteNTasks func(int) bool
    left, right := 0, min(len(tasks), len(workers))
    res := 0
    sort.Ints(tasks)
    sort.Ints(workers)

    canCompleteNTasks = func(tasksToComplete int) bool {
        availableWorkers := append([]int(nil), workers[len(workers)-tasksToComplete:]...)
        pillsLeft := pills

        for i := tasksToComplete - 1; i >= 0; i-- {
            n := len(availableWorkers)
            if n > 0 && availableWorkers[n-1] >= tasks[i] {
                availableWorkers = availableWorkers[:n-1]
            } else {
                if pillsLeft == 0 {
                    return false
                }

                minCompletingWorkerIdx := sort.Search(n, func(j int) bool { return availableWorkers[j] >= tasks[i] - strength })
                if minCompletingWorkerIdx == n {
                    return false
                }

                availableWorkers = append(availableWorkers[:minCompletingWorkerIdx], availableWorkers[minCompletingWorkerIdx+1:]...)
                pillsLeft -= 1
            }
        }

        return true
    }

    for left <= right {
        middle := left + (right - left) / 2

        if canCompleteNTasks(middle) {
            res = middle
            left = middle + 1
        } else {
            right = middle - 1
        }
    }

    return res
}