func minimumOperations(nums []int) int {
    freq := make(map[int]int)
    n := len(nums)

    for i := n - 1; i >= 0; i-- {
        freq[nums[i]] += 1

        if freq[nums[i]] > 1 {
            a := 0

            if (i + 1) % 3 != 0 {
                a = 1
            }

            return (i + 1) / 3 + a
        }
    }
    
    return 0
}