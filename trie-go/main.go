package main

import (
	"fmt"
	"os"
	"strconv"
)

func runeToString(runeValue rune) string {
	return strconv.QuoteRune(runeValue)
}

func getCharacterFromMap(hashmap map[rune]*Node) string {
	for key := range hashmap {
		return runeToString(key)
	}

	return ""
}

type Node struct {
	Children    map[rune]*Node
	IsCompleted bool
}

func newNode() *Node {
	return &Node{
		Children:    make(map[rune]*Node),
		IsCompleted: false,
	}
}

func (node *Node) setChildren(character rune) {
	node.Children[character] = newNode()
}

func (node *Node) setEOF() {
	node.IsCompleted = true
}

func (node *Node) unsetEOF() {
	node.IsCompleted = false
}

type Trie struct {
	Root    *Node
	Current *Node
	// Previous *Node
}

func newTrie() *Trie {
	return &Trie{
		Root:    newNode(),
		Current: nil,
		// Previous: nil,
	}
}

func (trie *Trie) insert(text string) {
	for _, character := range text {
		_, ok := trie.Root.Children[character]

		// fmt.Println("Root/Current:", trie.Root, trie.Current)

		if !ok && trie.Current == nil {
			// fmt.Printf("Insert %s to the root.\n", runeToString(character))
			trie.Root.setChildren(character)
			trie.Current = trie.Root.Children[character]
			// trie.Previous = trie.Root
			continue
		}

		if trie.Current != nil {
			_, exists := trie.Current.Children[character]

			if !exists {
				// fmt.Printf("Insert %s to be child of %s.\n", runeToString(character), getCharacterFromMap(trie.Previous.Children))
				trie.Current.setChildren(character)
				// trie.Previous = trie.Current
			}
		}

		trie.Current = trie.Current.Children[character]
	}

	trie.Current.setEOF()
	trie.Current = trie.Root
	// trie.Previous = trie.Current
}

func (trie *Trie) search(text string) bool {
	var lastNode *Node

	for _, character := range text {
		node, ok := trie.Current.Children[character]

		if !ok {
			trie.Current = trie.Root
			return false
		}

		trie.Current = node
		lastNode = node
	}

	trie.Current = trie.Root
	return lastNode.IsCompleted
}

func (trie *Trie) remove(text string) {
	var lastNode *Node

	for _, character := range text {
		node, ok := trie.Current.Children[character]

		if !ok {
			trie.Current = trie.Root
			fmt.Fprintf(os.Stderr, "[Stop in the middle] %s does not exist.\n", text)
			return
		}

		trie.Current = node
		lastNode = node
	}

	trie.Current = trie.Root

	if lastNode.IsCompleted {
		lastNode.unsetEOF()
		fmt.Printf("Unlink %s from the trie.\n", text)
		return
	}

	fmt.Fprintf(os.Stderr, "%s does not exist.\n", text)
}

// TODO: write unit tests and do the remove method
func main() {
	trie := newTrie()

	trie.insert("Vitor")
	trie.insert("Viagem")
	trie.insert("Andrew")
	trie.insert("Banana")
	trie.insert("Andrewx")
	trie.insert("Liz")
	trie.insert("bbbbb")
	trie.insert("ccccc")
	trie.insert("ddddd")
	trie.insert("eeeee")
	trie.insert("exe2eae4e")

	fmt.Printf("Does %s exists? %t\n", "Vitor", trie.search("Vitor"))
	fmt.Printf("Does %s exists? %t\n", "Banana", trie.search("Banana"))
	fmt.Printf("Does %s exists? %t\n", "Viagem", trie.search("Viagem"))
	fmt.Printf("Does %s exists? %t\n", "Andrew", trie.search("Andrew"))
	fmt.Printf("Does %s exists? %t\n", "Andrewx", trie.search("Andrewx"))
	fmt.Printf("Does %s exists? %t\n", "Liz", trie.search("Liz"))
	fmt.Printf("Does %s exists? %t\n", "bbbbb", trie.search("bbbbb"))
	fmt.Printf("Does %s exists? %t\n", "ccccc", trie.search("ccccc"))
	fmt.Printf("Does %s exists? %t\n", "ddddd", trie.search("ddddd"))
	fmt.Printf("Does %s exists? %t\n", "eeeee", trie.search("eeeee"))
	fmt.Printf("Does %s exists? %t\n", "exe2eae4e", trie.search("exe2eae4e"))

	fmt.Println()

	fmt.Printf("Does %s exists? %t\n", "liz", trie.search("liz"))
	fmt.Printf("Does %s exists? %t\n", "Audrew", trie.search("Audrew"))
	fmt.Printf("Does %s exists? %t\n", "John", trie.search("John"))
	fmt.Printf("Does %s exists? %t\n", "Viarew", trie.search("Viarew"))
	fmt.Printf("Does %s exists? %t\n", "Vitorx", trie.search("Vitorx"))
	fmt.Printf("Does %s exists? %t\n", "Vitorr", trie.search("Vitorr"))
	fmt.Printf("Does %s exists? %t\n", "eeeee2", trie.search("eeeee2"))
	fmt.Printf("Does %s exists? %t\n", "ee", trie.search("ee"))
	fmt.Printf("Does %s exists? %t\n", "a", trie.search("a"))

	fmt.Println()

	trie.remove("Vitor")
	trie.remove("Banana")
	trie.remove("Viagem")
	trie.remove("Liz")
	trie.remove("liz")
	trie.remove("vitorx")
	trie.remove("Audrew")
	trie.remove("a")
	trie.remove("exe2eae4e")

	fmt.Println()

	fmt.Printf("Does %s exists? %t\n", "Vitor", trie.search("Vitor"))
	fmt.Printf("Does %s exists? %t\n", "Banana", trie.search("Banana"))
	fmt.Printf("Does %s exists? %t\n", "Viagem", trie.search("Viagem"))
	fmt.Printf("Does %s exists? %t\n", "Liz", trie.search("Liz"))
	fmt.Printf("Does %s exists? %t\n", "liz", trie.search("liz"))
	fmt.Printf("Does %s exists? %t\n", "vitorx", trie.search("vitorx"))
	fmt.Printf("Does %s exists? %t\n", "Audrew", trie.search("Audrew"))
	fmt.Printf("Does %s exists? %t\n", "exe2eae4e", trie.search("exe2eae4e"))
}
