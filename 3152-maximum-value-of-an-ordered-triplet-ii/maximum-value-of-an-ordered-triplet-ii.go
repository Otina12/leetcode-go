func maximumTripletValue(nums []int) int64 {
    var maxNum, maxDiff, maxTriplet int = 0, 0, 0

    for _, num := range nums {
        maxTriplet = max(maxTriplet, num * maxDiff)
        maxDiff = max(maxDiff, maxNum - num)
        maxNum = max(maxNum, num)
    }

    return int64(maxTriplet)
}

func max(num1 int, num2 int) int {
    if num1 > num2 {
        return num1
    }

    return num2
}