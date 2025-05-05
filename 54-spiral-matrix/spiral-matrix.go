func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    directions := [][]int{[]int{0,1}, []int{1,0}, []int{0,-1}, []int{-1,0}}
    directionSteps := []int{n-2, m-1, n-1, m-2}
    curDirection := 1
    res := []int{}

    for j := 0; j < n; j++ {
        res = append(res, matrix[0][j])
    }

    curI, curJ := 0, n-1

    for {
        if directionSteps[curDirection] <= 0 {
            break
        }

        for i := 0; i < directionSteps[curDirection]; i++ {
            curI += directions[curDirection][0]
            curJ += directions[curDirection][1]
            res = append(res, matrix[curI][curJ])
        }

        directionSteps[curDirection] -= 2
        curDirection = (curDirection + 1) % 4
    }

    return res
}