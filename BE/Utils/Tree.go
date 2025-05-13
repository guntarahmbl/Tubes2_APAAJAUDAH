package utils

import (
	"fmt"
	"strings"
)

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

func isTreeSame(tree1, tree2 *TreeNode) bool {
	// Jika keduanya nil, mereka sama
	if tree1 == nil && tree2 == nil {
		return true
	}

	// Jika salah satu nil, mereka tidak sama
	if tree1 == nil || tree2 == nil {
		return false
	}

	// Periksa Item1
	if len(tree1.Item1) != len(tree2.Item1) || !isMapEqual(tree1.Item1, tree2.Item1) {
		return false
	}

	// Periksa Item2
	if len(tree1.Item2) != len(tree2.Item2) || !isMapEqual(tree1.Item2, tree2.Item2) {
		return false
	}

	// Periksa jumlah anak di Children1 dan Children2
	if len(tree1.Children1) != len(tree2.Children1) || len(tree1.Children2) != len(tree2.Children2) {
		return false
	}

	// Rekursi untuk setiap anak di Children1
	for i := 0; i < len(tree1.Children1); i++ {
		if !isTreeSame(tree1.Children1[i], tree2.Children1[i]) {
			return false
		}
	}

	// Rekursi untuk setiap anak di Children2
	for i := 0; i < len(tree1.Children2); i++ {
		if !isTreeSame(tree1.Children2[i], tree2.Children2[i]) {
			return false
		}
	}

	// Jika semua pemeriksaan lulus, pohon identik
	return true
}

// Fungsi pembantu untuk membandingkan dua map
func isMapEqual(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val1 := range map1 {
		val2, exists := map2[key]
		if !exists || val1 != val2 {
			return false
		}
	}

	return true
}
