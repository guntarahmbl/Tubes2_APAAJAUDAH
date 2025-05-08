package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	// "time"
)

var memo = make(map[string]*TreeNode)

type TreeNode struct {
	Item1    string
	Item2    string
	Children []*TreeNode
}


func isBaseElement(item string) bool {
	switch item {
	case "Water", "Fire", "Air", "Earth":
		return true
	}
	return false
}

func isLeaf(node TreeNode) bool {
	if node.Item2 == "" {
		return true
	}
	return isBaseElement(node.Item1) && isBaseElement(node.Item2)
}

func buildTreeBFS(target string, result map[string][][]string, root *TreeNode) {
	// Antrian untuk BFS
	queue := []*TreeNode{root}

	// Proses BFS
	for len(queue) > 0 {
		// Ambil node pertama dalam antrian
		currentNode := queue[0]
		queue = queue[1:] // Dequeue

		// iterasi item1
		if _, found := memo[currentNode.Item1]; !found {
			// Cari resep untuk target node saat ini
			recipes1, ok := result[currentNode.Item1]
			if !ok {
				continue
			}
			// Iterasi melalui semua pasangan resep1
			for _, val := range recipes1 {
				newNode := &TreeNode{
					Item1: val[0],
					Item2: val[1],
				}
				// time.Sleep(time.Second)
				// fmt.Printf("%s : %s - %s\n", currentNode.Item1, val[0], val[1])
				
				if (isLeaf(*newNode)) {
					continue;
				}
	
				currentNode.Children = append(currentNode.Children, newNode)
	
				// mencegah loop
				if (val[0] != target && val[1] != target) {
					queue = append(queue, newNode)
				}
	
			}
			memo[currentNode.Item1] = currentNode
		}

		// iterasi item2
		if _, found := memo[currentNode.Item2]; !found {
			// Cari resep untuk target node saat ini
			recipes2, ok := result[currentNode.Item2]
			if !ok {
				continue
			}
			// Iterasi melalui semua pasangan resep1
			for _, val := range recipes2 {
				newNode := &TreeNode{
					Item1: val[0],
					Item2: val[1],
				}
				// time.Sleep(time.Second)
				// fmt.Printf("%s : %s - %s\n", currentNode.Item1, val[0], val[1])
				
				if (isLeaf(*newNode)) {
					continue;
				}
	
				currentNode.Children = append(currentNode.Children, newNode)
	
				// mencegah loop
				if (val[0] != target && val[1] != target) {
					queue = append(queue, newNode)
				}
	
			}
			memo[currentNode.Item2] = currentNode
		}
	}
}

func traverseTree(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	fmt.Printf("%s%s, %s\n", strings.Repeat("-", depth), node.Item1, node.Item2)
	for _, child := range node.Children {
		traverseTree(child, depth+1)
	}
}

func main() {
	// Buka dan baca JSON
	f, err := os.Open("../data/elements_recipes_real.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var result map[string][][]string
	if err := json.Unmarshal(b, &result); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	root := &TreeNode{Item1: "Pressure"}
	buildTreeBFS("Pressure", result, root)
	traverseTree(root, 0)
}
