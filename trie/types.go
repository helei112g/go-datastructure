package trie

const SIZE = 26

type TrieNode struct {
	data      rune
	children  [SIZE]*TrieNode
	isEndChar bool
}
