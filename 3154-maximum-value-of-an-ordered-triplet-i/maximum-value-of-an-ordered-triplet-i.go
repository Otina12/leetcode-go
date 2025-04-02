func maximumTripletValue(nums []int) int64 {
    var n int = len(nums)
    var res int = 0

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                res = max(res, (nums[i] - nums[j]) * nums[k])
            }
        }
    }

    return int64(res)
}

func max(num1 int, num2 int) int {
    if num1 > num2 {
        return num1
    }

    return num2
}