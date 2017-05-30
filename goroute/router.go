package goroute

func sdf() {
	trie := new(node)
	trie.Insert("armin/bob")
	trie.Insert("armin/bob/bill")
	trie.Insert("armin/bob/xun")
	trie.Insert("armin/xun")
	trie.Insert("armin/xun/bob")
	trie.Insert("xun/bob")
	trie.Insert("a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z")
	trie.PrintTrie("")
}


