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

var memo = make(map[string][]*TreeNode)

type TreeNode struct {
	Item1    string
	Item2    string
	Children1 []*TreeNode
	Children2 []*TreeNode
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
func displayTree(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	// Indentasi berdasarkan depth
	prefix := strings.Repeat("-", depth)

	// Cetak node saat ini
	fmt.Printf("%s%s (%s)\n", prefix, node.Item1, node.Item2)

	// Tampilkan Children1
	for _, child := range node.Children1 {
		displayTree(child, depth+1)
	}

	// Tampilkan Children2
	for _, child := range node.Children2 {
		displayTree(child, depth+1)
	}
}

func buildTreeBFS(result map[string][][]string, root *TreeNode, tier map[string]int) {
	// Queue untuk BFS
	queue := []*TreeNode{root}

	// Proses BFS
	for len(queue) > 0 {
		
		currentNode := queue[0] // Ambil node pertama dalam queue
		queue = queue[1:] // Dequeue


		// iterasi item1
		if memoChildren, found := memo[currentNode.Item1]; !found { // cari di memo

			recipes1 := result[currentNode.Item1]  // cari resep-resep item1

			for _, val := range recipes1 {  // iterasi tiap resep

				// jika tier sama atau dibawah, jangan dianmbil resepnya 
				if (tier[currentNode.Item1] <= tier[val[0]] || tier[currentNode.Item1] <= tier[val[1]]) {
					continue;
				}

				newNode := &TreeNode{ // buat node baru
					Item1: val[0],
					Item2: val[1],
				}
				currentNode.Children1 = append(currentNode.Children1, newNode) // masukkan sebagai children1
				
				queue = append(queue, newNode) // masukkan ke queue
	
			}

			memo[currentNode.Item1] = currentNode.Children1 // catat bahwa Item1 sudah pernah di proses
		} else {
			currentNode.Children1 = append(currentNode.Children1, memoChildren...) // ambil dari memo jika Item1 sudah pernah di proses
		}

		// iterasi item2
		if memoChildren, found := memo[currentNode.Item2]; !found {

			recipes2 := result[currentNode.Item2]

			for _, val := range recipes2 {

				if (tier[currentNode.Item2] <= tier[val[0]] || tier[currentNode.Item2] <= tier[val[1]]) {
					continue;
				}

				newNode := &TreeNode{
					Item1: val[0],
					Item2: val[1],
				}
				currentNode.Children2 = append(currentNode.Children2, newNode)
				
				queue = append(queue, newNode)
			}
			memo[currentNode.Item2] = currentNode.Children2
		} else {
			currentNode.Children2 = append(currentNode.Children2, memoChildren...)
		}
	}
}

func traverseTreeDFS(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	// Cetak informasi node saat ini
	fmt.Printf("%s %s, %s \n", strings.Repeat("-", depth), node.Item1, node.Item2)

	// Rekursi untuk Children1
	for _, child := range node.Children1 {
		traverseTreeDFS(child, depth+1)
	}

	// Rekursi untuk Children2
	for _, child := range node.Children2 {
		traverseTreeDFS(child, depth+1)
	}
}

func traverseTreeBFS(root *TreeNode) {
	if root == nil {
		return
	}

	// Queue untuk BFS
	type NodeWithDepth struct {
		node  *TreeNode
		depth int
		parent string
	}

	queue := []NodeWithDepth{{node: root, depth: 0, parent: ""}}

	for len(queue) > 0 {
		// Ambil node pertama dari antrian
		current := queue[0]
		queue = queue[1:]

		// Cetak informasi node saat ini
		fmt.Printf("%s %s, %s \n", strings.Repeat("-", current.depth), current.node.Item1, current.node.Item2)

		// Tambahkan Children1 ke antrian
		for _, child := range current.node.Children1 {
			queue = append(queue, NodeWithDepth{node: child, depth: current.depth + 1, parent: current.node.Item1})
		}

		// Tambahkan Children2 ke antrian
		for _, child := range current.node.Children2 {
			queue = append(queue, NodeWithDepth{node: child, depth: current.depth + 1, parent: current.node.Item2})
		}
	}
}


func readRecipes(filePath string) (map[string][][]string, error) {
	// Buka file JSON
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Baca
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Parse JSON ke dalam map
	var result map[string][][]string
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func readElementsTier(filePath string) (map[string]int, error) {
	// Buka file JSON
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Baca
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parsing JSON ke slice map sementara karena gabisa langsung map[string]int
	var tempElements []map[string]interface{}
	if err := json.Unmarshal(data, &tempElements); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Konversi ke map[string]int
	elements := make(map[string]int)
	for _, elem := range tempElements {
		name := elem["nama"].(string)
		tier := int(elem["tier"].(float64)) 
		elements[name] = tier
	}

	return elements, nil
}


func main() {

	// Memanggil readRecipes
	recipes, err := readRecipes("../data/elements_recipes_real.json")
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}
	// Memanggil readElementsTier
	tier, err := readElementsTier("../data/allElementsTiers.json")
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	root := &TreeNode{Item1: "Rust"}
	buildTreeBFS(recipes, root, tier)
	traverseTreeDFS(root, 0)
}
