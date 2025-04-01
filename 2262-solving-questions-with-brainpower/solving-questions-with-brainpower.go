var memo []int64
var _questions [][]int

func mostPoints(questions [][]int) int64 {
    memo = make([]int64, len(questions))
    _questions = questions
    helper(0)

    return memo[0]
}

func helper(i int) int64 {
    if i >= len(_questions) {
        return 0
    }

    if memo[i] != 0 {
        return memo[i]
    }

    var skip int64 = helper(i + 1)
    var take int64 = int64(_questions[i][0]) + helper(i + _questions[i][1] + 1)
    memo[i] = max(skip, take)

    return memo[i]
}

func max(num1 int64, num2 int64) int64 {
    if num1 > num2 {
        return num1
    }

    return num2
}