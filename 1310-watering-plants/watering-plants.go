func wateringPlants(plants []int, capacity int) int {
    steps := 0
    curCapacity := capacity
    i := 0

    for i < len(plants) {
        if curCapacity < plants[i] {
            curCapacity = capacity
            steps += 2 * i
        }

        curCapacity -= plants[i]
        steps += 1
        i += 1
    }

    return steps
}