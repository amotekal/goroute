package goroute

import (
	"fmt"
	"net/http"
	"strings"
)

type nodeType int

const (
	static nodeType = iota
	param
	wild
	catchAll
)

//node represents a level in the trie
type node struct {
	path string

	children map[string]*node

	wildChild *node

	nodeType

	handle *http.HandlerFunc
}

//Insert a new path into the trie
func (n *node) insert(path string) error {
	parts := strings.Split(path, "/")
	curr := n

	for _, part := range parts {
		if n.path == "" {
			n.path = part
			n.nodeType = static
		} else if next, contains := curr.children[part]; contains {
			curr = next
		} else {
			newNode := new(node)
			curr.children[part] = newNode
			curr = newNode
		}
	}
	return nil
}

//PrintTrie prints the trie to standard out
func (n *node) printTrie(depth string) {
	fmt.Println(depth + n.path)
	for _, next := range n.children {
		next.printTrie(depth + " ")
	}
}
