func minImpossibleOR(nums []int) int {
    numSet := make(map[int]struct{})

    for _, num := range nums {
        if num & (num - 1) == 0 {
            numSet[num] = struct{}{}
        }
    }

    x := 1

    for {
        if _, exists := numSet[x]; !exists {
            break
        } else {
            x <<= 1
        }
    }

    return x
}