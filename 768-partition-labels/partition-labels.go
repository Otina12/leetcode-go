func partitionLabels(s string) []int {
    var occurences [][]int = make([][]int, 26)

    for i := 0; i < 26; i++ {
        occurences[i] = []int {0, 0}
    } 

    for i, letter := range s {
        var idx = int(letter) - 97
        if occurences[idx][0] == 0 {
            occurences[idx][0] = i
        }
        occurences[idx][1] = i
    }

    var res []int = []int {}
    var lastTakeIdx int = 0
    var right int = 0

    for i, letter := range s {
        var idx = int(letter) - 97
        if occurences[idx][1] > right {
            right = occurences[idx][1]
        }
        if right == i {
            res = append(res, right - lastTakeIdx + 1)
            lastTakeIdx = i + 1
        }
    }

    return res
}