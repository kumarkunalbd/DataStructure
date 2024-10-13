package trie

import (
	"fmt"
)

func prefix(A []string) []string {
	root := newTrie()

	// tc : O(n*m)
	for _, word := range A {
		root.insert(word)
	}
	ufs := make([]string, 0)

	for _, word := range A {
		uf := root.serachuniquep(word)
		ufs = append(ufs, uf)
	}

	return ufs
}

func (tr *Trie) serachuniquep(word string) string {

	cr := tr.charnode[word[0]]
	count := 0
	for i := 1; i < len(word); i++ {
		count = i
		if len(cr.charnode) > 1 {
			ch := word[i]
			fmt.Printf("ch=%v,%c\n", ch)

			cr = cr.charnode[ch]
			continue
		}

		break
	}
	return word[:count]
}

func (tr *Trie) insert(word string) {
	cur := tr
	for i := 0; i < len(word); i++ {
		at := cur.charnode[word[i]]
		if at == nil {
			at = newTrie()
			cur.charnode[word[i]] = at
		}
		cur = at
	}
	cur.isend = true
}

type Trie struct {
	isend    bool
	charnode map[byte]*Trie
}

func newTrie() *Trie {
	tri := Trie{
		isend:    false,
		charnode: make(map[byte]*Trie, 0),
	}
	return &tri
}
