func numTriplets(nums1 []int, nums2 []int) int {
    return countTriplets(nums1, nums2) + countTriplets(nums2, nums1)
}

func countTriplets(squares []int, products []int) int {
    freq := make(map[int]int)

    for _, num := range products {
        freq[num] += 1
    }

    res := 0

    for _, a := range squares {
        target := a * a

        for b, freqB := range freq {
            if target % b != 0 {
                continue
            }

            c := target / b
            if b > c {
                continue
            }
            
            if freqC, exists := freq[c]; exists {
                if b == c {
                    res += freqB * (freqB - 1) / 2
                } else {
                    res += freqB * freqC
                }
            }
        }
    }

    return res
}