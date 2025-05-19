func lexicalOrder(n int) []int {
    res := make([]int, n)
    stack := []int{1}
    res[0] = 1
    curIdx := 0

    for len(stack) > 0 {
        num := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if num > n {
            continue
        }

        res[curIdx] = num
        curIdx += 1
        
        if num % 10 != 9 {
            stack = append(stack, num + 1)
        }
        stack = append(stack, num * 10)
    }

    return res
}