func numRabbits(answers []int) int {
    freq := make(map[int]int)

    for _, answer := range answers {
        freq[answer] += 1
    }

    res := 0
    for key, value := range freq {
        res += ceilingDivision(value, key + 1) * (key + 1)
    }

    return res
}

func ceilingDivision(num1, num2 int) int {
    return (num1 + num2 - 1) / num2
}