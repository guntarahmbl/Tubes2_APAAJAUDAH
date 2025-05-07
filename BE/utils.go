package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TreeNode struct {
	item1    string
	item2    string
	children []*TreeNode
}

func isBaseElement(item string) bool {
	switch item {
	case "Water", "Fire", "Air", "Earth":
		return true
	}
	return false
}

func isLeaf(node TreeNode) bool {
	if node.item2 == "" {
		return true
	}
	return isBaseElement(node.item1) && isBaseElement(node.item2)
}

func buildTree(target string, result map[string][][]string, node *TreeNode) {
    // base: no recipes for this target, atau sudah leaf
    recipes, ok := result[target]
    if !ok || isLeaf(*node) {
        return
    }

    for _, val := range recipes {
        // buat node baru untuk setiap pasangan
        newNode := &TreeNode{
            item1: val[0],
            item2: val[1],
        }
        // rekursi pada masing‑masing bahan
        buildTree(val[0], result, newNode)
        buildTree(val[1], result, newNode)

        node.children = append(node.children, newNode)
    }
}

func traverseTree(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	indent := strings.Repeat("  ", depth)
	if len(node.children) == 0 {
		fmt.Printf("%sLeaf: %s + %s\n", indent, node.item1, node.item2)
		return
	}
	fmt.Printf("%sNode: %s + %s\n", indent, node.item1, node.item2)
	for _, child := range node.children {
		traverseTree(child, depth+1)
	}
}

func main() {
	// buka dan baca JSON
	f, err := os.Open("data/elements_recipes_real.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][][]string
	if err := json.Unmarshal(b, &result); err != nil {
		log.Fatal(err)
	}

	seen := make(map[string]bool)
	cache := make(map[string]*TreeNode)

	root := &TreeNode{item1: "Brick", item2: "Brick"}
	buildTree("Brick", &root, result, seen, cache)
	traverseTree(root, 0)
	// for _,child := range root.children {
	// 	fmt.Print(child.item1, child.item2)
	// }
}
