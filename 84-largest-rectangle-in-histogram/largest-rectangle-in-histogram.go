type pair struct {
    index int
    value int 
}

func largestRectangleArea(heights []int) int {
    monoStack := []pair{}
    maxArea := 0

    for i, height := range heights {
        if len(monoStack) > 0 && monoStack[len(monoStack)-1].value >= height {
            poppedBarStartIndex := -1

            for len(monoStack) > 0 && monoStack[len(monoStack)-1].value >= height {
                poppedBar := monoStack[len(monoStack)-1]
                poppedBarStartIndex = poppedBar.index
                maxArea = max(maxArea, (i - poppedBarStartIndex) * poppedBar.value)
                monoStack = monoStack[:len(monoStack)-1]
            }

            monoStack = append(monoStack, pair{poppedBarStartIndex, height})
        } else {
            monoStack = append(monoStack, pair{i, height})
        }
    }

    for len(monoStack) > 0 {
        lastBar := monoStack[len(monoStack)-1]
        maxArea = max(maxArea, (len(heights) - lastBar.index) * lastBar.value)
        monoStack = monoStack[:len(monoStack)-1]
    }

    return maxArea
}