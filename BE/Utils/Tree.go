package utils

import (
	"fmt"
	"strings"
)

// menerapkan memoization, data yang pernah dicari tidak usah dicari lagi
var memo = make(map[string][]*TreeNode) 

// struktur tree
type TreeNode struct {
	Item1    map[string]string
	Item2    map[string]string
	Children1 []*TreeNode
	Children2 []*TreeNode
}

// variable NullChild
var NullChild = &TreeNode{
		Item1: map[string]string{"Name":  "Null"},
		Item2: map[string]string{"Name":  "Null"},
}

// checker
func isBaseElement(item string) bool {
	switch item {
	case "Water", "Fire", "Air", "Earth":
		return true
	}
	return false
}
func isLeaf(node TreeNode) bool {
	if node.Item2["Name"] == "" {
		return true
	}
	return isBaseElement(node.Item1["Name"]) && isBaseElement(node.Item2["Name"])
}

// menampilkan tree (secara DFS)
func PrintTree(node *TreeNode, level int) {

	if node == nil {
		return
	}
	indent := strings.Repeat("- ", level)

	// Print the current node's item
	fmt.Printf("%s%s, %s\n", indent, node.Item1["Name"], node.Item2["Name"])

	// Recursively print the children
	if len(node.Children1) > 0 {
		PrintTree(node.Children1[0], level+1)
	}
	if len(node.Children2) > 0 {
		PrintTree(node.Children2[0], level+1)
	}
}

// manampilkan tree-tree pada suatu list
func PrintListOfTree(recipes []*TreeNode) {
	for i,recipe := range recipes {
		fmt.Printf("==========================\n")
		fmt.Printf("Tree %d:\n", i+1)
		PrintTree(recipe, 0)
	}
	fmt.Printf("==========================\n")
}
