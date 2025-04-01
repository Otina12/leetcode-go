type Pair struct {
	X int
	Y int
}

var _word1 string
var _word2 string
var memo map[Pair]int

func minDistance(word1 string, word2 string) int {
    _word1 = word1
    _word2 = word2
    memo = make(map[Pair]int)

    return helper(0, 0)
}

func helper(i int, j int) int {
    if i == len(_word1) {
        return len(_word2) - j
    }

    if j == len(_word2) {
        return len(_word1) - i
    }

    if val, ok := memo[Pair{i, j}]; ok {
        return val
    }
    
    if _word1[i] == _word2[j] {
        return helper(i + 1, j + 1)
    }

    var insert int = 1 + helper(i, j + 1) // insert word2[j] into word1, so j is incremented
    var delete int = 1 + helper(i + 1, j) // delete word1[i], so i is incremented
    var replace int = 1 + helper(i + 1, j + 1) // replace word1[i] with word2[j], so both are incremented

    var res int = min(insert, min(delete, replace))
    memo[Pair{i, j}] = res
    
    return res
}

func min(num1 int, num2 int) int {
    if num1 < num2 {
        return num1
    }

    return num2
}