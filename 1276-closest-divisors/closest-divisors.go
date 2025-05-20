func closestDivisors(num int) []int {
    num1, num2 := num + 1, num + 2
    upperbound := math.Sqrt(float64(num2))
    minDiff := 1_000_000_002
    res := []int{0, 0}

    for x := 1; x <= int(upperbound) + 1; x++ {
        if num1 % x == 0 {
            div := num1 / x
            diff := abs(div - x)
            if diff < minDiff {
                res[0], res[1] = x, div
                minDiff = diff
            }
        }
        if num2 % x == 0 {
            div := num2 / x
            diff := abs(div - x)
            if diff < minDiff {
                res[0], res[1] = x, div
                minDiff = diff
            }
        }
    }

    return res
}

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}