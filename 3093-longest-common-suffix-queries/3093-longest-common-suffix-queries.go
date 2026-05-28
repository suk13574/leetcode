type TrieNode struct {
	children [26]*TrieNode
	bestIdx  int
	bestLen  int
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		bestIdx: -1,
		bestLen: 1 << 30,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string, idx int) {
	node := t.root
	n := len(word)

	updateBest(node, n, idx)

	for i := n - 1; i >= 0; i-- {
		c := int(word[i] - 'a')

		if node.children[c] == nil {
			node.children[c] = NewTrieNode()
		}

		node = node.children[c]
		updateBest(node, n, idx)
	}
}

func (t *Trie) Query(word string) int {
	node := t.root
	n := len(word)

	for i := n - 1; i >= 0; i-- {
		c := int(word[i] - 'a')

		if node.children[c] == nil {
			break
		}

		node = node.children[c]
	}

	return node.bestIdx
}

func better(length, idx, bestLen, bestIdx int) bool {
	if bestIdx == -1 {
		return true
	}

	if length < bestLen {
		return true
	}

	if length == bestLen && idx < bestIdx {
		return true
	}

	return false
}

func updateBest(node *TrieNode, length int, idx int) {
	if better(length, idx, node.bestLen, node.bestIdx) {
		node.bestLen = length
		node.bestIdx = idx
	}
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	trie := NewTrie()

	for i, w := range wordsContainer {
		trie.Insert(w, i)
	}

	res := make([]int, len(wordsQuery))

	for i, w := range wordsQuery {
		res[i] = trie.Query(w)
	}

	return res
}