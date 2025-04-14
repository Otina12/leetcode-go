func countGoodTriplets(arr []int, a int, b int, c int) int {
    res := 0
    n := len(arr)
    
    for i := 0; i < n - 2; i++ {
        for j := i + 1; j < n - 1; j++ {
            if abs(arr[i] - arr[j]) > a {
                continue
            }
            for k := j+1; k < n; k++ {
                if abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
                    res += 1
                }
            }
        }
    }
    
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}