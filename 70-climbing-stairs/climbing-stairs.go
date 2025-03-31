func climbStairs(n int) int {
    if n <= 2 {
        return n;
    }

    var prevByOne int = 2
    var prevByTwo int = 1
    var res int = 0

    for i := 3; i <= n; i++ {
        res = prevByTwo + prevByOne
        prevByTwo = prevByOne
        prevByOne = res
    }

    return res
}