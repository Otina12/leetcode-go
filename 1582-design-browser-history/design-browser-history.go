type BrowserHistory struct {
    pages []string
    currentPageIndex int
    lastPageIndex int
}


func Constructor(homepage string) BrowserHistory {
    pages := []string { homepage }
    
    return BrowserHistory { pages, 0, 0 }
}


func (this *BrowserHistory) Visit(url string)  {
    if this.currentPageIndex + 2 > len(this.pages) {
        this.pages = grow(this.pages)
    }

    this.pages[this.currentPageIndex + 1] = url
    this.currentPageIndex += 1
    this.lastPageIndex = this.currentPageIndex
}


func (this *BrowserHistory) Back(steps int) string {
    steps = min(steps, this.currentPageIndex)
    page := this.pages[this.currentPageIndex - steps]
    this.currentPageIndex -= steps

    return page
}


func (this *BrowserHistory) Forward(steps int) string {
    steps = min(steps, this.lastPageIndex - this.currentPageIndex)
    page := this.pages[this.currentPageIndex + steps]
    this.currentPageIndex += steps

    return page
}

func grow(arr []string) []string {
    newLen := max(1, len(arr) * 2)
    resizedArr := make([]string, newLen)
    copy(resizedArr, arr)

    return resizedArr
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */