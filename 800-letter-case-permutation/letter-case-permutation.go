func letterCasePermutation(s string) []string {
    curString := strings.Split(s, "")
    var res []string
    var backtrack func(int)

    backtrack = func(i int) {
        if i == len(s) {
            res = append(res, strings.Join(curString, ""))
            return
        }

        backtrack(i + 1)

        if isDigit(s[i]) {
            originalCharacter := curString[i]
            curString[i] = getOthercaseCharacter(originalCharacter)
            backtrack(i + 1)
            curString[i] = originalCharacter
        }
    }

    backtrack(0)
    return res
}

func isDigit(character byte) bool {
    return character >= 'A' && character <= 'Z' || character >= 'a' && character <= 'z'
}

func getOthercaseCharacter(characterStr string) string {
    character := characterStr[0]

    if character >= 'A' && character <= 'Z' {
        return string(character + 32)
    } else {
        return string(character - 32)
    }
}