func countAndSay(n int) string {
    var curSequence string = "1"

    for i := 1; i < n; i++ {
        curSequence = nextSequence(curSequence)
    }

    return curSequence
}

func nextSequence(sequence string) string {
	runes := []rune(sequence)
	curChar := runes[0]
	curCnt := 1
	var sb strings.Builder

	for i := 1; i < len(runes); i++ {
		if runes[i] != curChar {
			sb.WriteString(strconv.Itoa(curCnt))
			sb.WriteRune(curChar)
			curChar = runes[i]
			curCnt = 0
		}
		
        curCnt += 1
	}

	sb.WriteString(strconv.Itoa(curCnt))
	sb.WriteRune(curChar)

	return sb.String()
}