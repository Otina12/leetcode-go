func countDistinctIntegers(nums []int) int {
    var seen = make(map[int]bool)

    for _, num := range nums {
        seen[num] = true
        seen[reverseInteger(num)] = true
    }

    return len(seen)
}

func reverseInteger(num int) int {
    var res = 0

    for num > 0 {
        var digit = num % 10
        res = res * 10 + digit
        num /= 10
    }

    return res
}