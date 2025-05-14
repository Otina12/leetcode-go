func maximumTastiness(price []int, k int) int {
    sort.Ints(price)
    var isValidBasketTastiness func(int) bool

    isValidBasketTastiness = func(tastiness int) bool {
        cnt := 1
        prevChosenCandy := price[0]

        for i := 1; i < len(price); i++ {
            if price[i] - prevChosenCandy >= tastiness {
                cnt += 1
                prevChosenCandy = price[i]
            }

            if cnt == k {
                return true
            }
        }

        return false
    }

    left, right := 0, price[len(price)-1] - price[0]
    res := 0

    for left <= right {
        middle := left + (right - left) / 2

        if isValidBasketTastiness(middle) {
            res = max(res, middle)
            left = middle + 1
        } else {
            right = middle - 1
        }
    }

    return res
}