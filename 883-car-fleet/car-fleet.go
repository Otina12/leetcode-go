type Car struct {
    Position int
    Speed int
}

func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    cars := make([]Car, n)

    for i := 0; i < n; i++ {
        cars[i] = Car{position[i], speed[i]}
    }

    sort.Slice(cars, func(i int, j int) bool {
        return cars[i].Position < cars[j].Position
    })

    var monoStack []float32

    for i := 0; i < n; i++ {
        timeToReach := float32(target - cars[i].Position) / float32(cars[i].Speed)

        for len(monoStack) > 0 && monoStack[len(monoStack) - 1] <= timeToReach {
            monoStack = monoStack[:len(monoStack) - 1]
        }

        monoStack = append(monoStack, timeToReach)
    }

    return len(monoStack)
}