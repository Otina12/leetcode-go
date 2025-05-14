func maxPalindromesAfterOperations(words []string) int {
    wordLengths := []int{}
    letterFreq := make([]int, 26)

    for _, word := range words {
        wordLengths = append(wordLengths, len(word))

        for _, letter := range word {
            letterFreq[int(letter - 'a')] += 1
        }
    }

    sort.Ints(wordLengths)
    res := 0

    for i := 0; i < len(wordLengths); i++ {
        charactersNeeded := wordLengths[i]

        if charactersNeeded & 1 == 1 {
            for j := 0; j < 26; j++ {
                if letterFreq[j] & 1 == 1 {
                    letterFreq[j] -= 1
                    charactersNeeded -= 1 // charactersNeeded becomes even
                    break
                }
            }
        }

        for j := 0; j < 26; j++ {
            if letterFreq[j] == 1 { // never leave a single odd letter
                continue
            }

            if letterFreq[j] >= charactersNeeded {
                letterFreq[j] -= charactersNeeded
                res += 1
                break
            } else { // only decrement by even count
                charactersNeeded -= letterFreq[j] - letterFreq[j] & 1
                letterFreq[j] = letterFreq[j] & 1 
            }
        }
    }

    return res
}