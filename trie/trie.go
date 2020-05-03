package trie

// 插入字符串
func (t *TrieNode) Insert(test string) (ok bool) {
	if t == nil {
		return false
	}

	// 插入
	p := t
	for _, ch := range test {
		index := ch - rune('a')

		// 不在其中
		if p.children[index] == nil {
			newNode := &TrieNode{
				data:      ch,
				children:  [26]*TrieNode{},
				isEndChar: false,
			}
			p.children[index] = newNode
		}
		p = p.children[index]
	}
	p.isEndChar = true

	return true
}

// 查找是否存在某个字符串
func (t *TrieNode) Find(pattern string) bool {
	if t == nil {
		return false
	}
	p := t
	for _, ch := range pattern {
		// 如果匹配，持续找下去
		index := ch - rune('a')
		if p.children[index] == nil {
			return false
		}

		p = p.children[index]
	}

	if p.isEndChar == true {
		return true
	}

	return true
}
