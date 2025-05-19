func countLatticePoints(circles [][]int) int {
    res := 0

    for x := 0; x <= 200; x++ {
        for y := 0; y <= 200; y++ {
            for _, circle := range circles {
                if isInsideCircle([]int{x,y}, circle) {
                    res += 1
                    break
                }
            }
        }
    }
    
    return res
}

func isInsideCircle(point []int, circle []int) bool {
    dx := float64(circle[0] - point[0])
    dy := float64(circle[1] - point[1])
    distance := math.Sqrt(dx*dx + dy*dy)
    return distance <= float64(circle[2])
}