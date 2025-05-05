type WordDictionary struct {
    root *Trie
}


func Constructor() WordDictionary {
    return WordDictionary{NewTrie()}
}

func (this *WordDictionary) AddWord(word string)  {
    this.root.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
    return this.root.Matches(word)
}

// Trie
type Trie struct {
    Children []*Trie
    IsLeaf bool
}

func NewTrie() *Trie {
    return &Trie{Children: make([]*Trie, 26, 26)}
}

func (this *Trie) Insert(word string) {
    curNode := this

    for _, ch := range word {
        charIdx := int(ch) - 97
        if curNode.Children[charIdx] == nil {
            curNode.Children[charIdx] = &Trie{Children: make([]*Trie, 26, 26)}
        }
        curNode = curNode.Children[charIdx]
    }

    curNode.IsLeaf = true
}

func (this *Trie) Matches(pattern string) bool {
    curNodes := []*Trie{this}

    for _, ch := range pattern {
        newNodes := []*Trie{}

        if ch == '.' {
            for _, node := range curNodes {
                for _, child := range node.Children {
                    if child != nil {
                        newNodes = append(newNodes, child)
                    }
                }
            }
        } else {
            charIdx := int(ch) - 97
            for _, node := range curNodes {
                if node.Children[charIdx] != nil {
                    newNodes = append(newNodes, node.Children[charIdx])
                }
            }
        }

        curNodes = newNodes
    }

    for _, node := range curNodes {
        if node.IsLeaf {
            return true
        }
    }

    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */