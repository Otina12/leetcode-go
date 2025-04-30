func findNumbers(nums []int) int {
    res := 0

    for _, num := range nums {
        if hasEvenNumberOfDigits(num) {
            res += 1
        }
    }

    return res
}

func hasEvenNumberOfDigits(num int) bool {
    hasEven := true

    for num > 0 {
        num /= 10
        hasEven = !hasEven
    }

    return hasEven
}