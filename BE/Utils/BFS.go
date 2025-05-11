package utils

import (
	"fmt"
	"strings"
)

// membangun pohon dengan algoritma BFS
func BuildTreeBFS(result map[string][][]string, root *TreeNode, tier map[string]int) {
	// Queue untuk BFS
	queue := []*TreeNode{root}

	// Proses BFS
	for len(queue) > 0 {
		
		currentNode := queue[0] // Ambil node pertama dalam queue
		queue = queue[1:] // Dequeue

		// iterasi item1
		name1 := currentNode.Item1["Name"]
		if memoChildren, found := memo[name1]; !found { // jika tidak ada di memo (belum pernah di proses)


			recipes1 := result[name1]  // cari resep-resep item1

			for _, val := range recipes1 {  // iterasi tiap resep

				
				// jika tier sama atau dibawah, jangan dianmbil resepnya 
				if (tier[name1] <= tier[val[0]] || tier[name1] <= tier[val[1]]) {
					continue;
				}

				newNode := &TreeNode{
					Item1: map[string]string{
						"Name":  val[0],
						"Image": fmt.Sprintf("images/%s.png", val[0]),
					},
					Item2: map[string]string{
						"Name":  val[1],
						"Image": fmt.Sprintf("images/%s.png", val[1]),
					},
				}
				currentNode.Children1 = append(currentNode.Children1, newNode) // masukkan sebagai children1
				
				queue = append(queue, newNode) // masukkan ke queue
	
			}

			memo[name1] = currentNode.Children1 // catat bahwa Item1 sudah pernah di proses
		} else {
			currentNode.Children1 = append(currentNode.Children1, memoChildren...) // ambil dari memo jika Item1 sudah pernah di proses
		}

		// iterasi item2
		name2 := currentNode.Item2["Name"]
		if memoChildren, found := memo[name2]; !found {

			recipes2 := result[name2]

			for _, val := range recipes2 {

				if (tier[name2] <= tier[val[0]] || tier[name2] <= tier[val[1]]) {
					continue;
				}

				newNode := &TreeNode{
					Item1: map[string]string{
						"Name":  val[0],
						"Image": fmt.Sprintf("images/%s.png", val[0]),
					},
					Item2: map[string]string{
						"Name":  val[1],
						"Image": fmt.Sprintf("images/%s.png", val[1]),
					},
				}

				currentNode.Children2 = append(currentNode.Children2, newNode)
				
				queue = append(queue, newNode)
			}
			memo[name2] = currentNode.Children2
		} else {
			currentNode.Children2 = append(currentNode.Children2, memoChildren...)
		}
	}
}

// melakukan traverse dengan algoritma BFS
func TraverseTreeBFS(root *TreeNode) {
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
		fmt.Printf("%s %s, %s \n", strings.Repeat("-", current.depth), current.node.Item1["Name"], current.node.Item2["Name"])

		// Tambahkan Children1 ke antrian
		for _, child := range current.node.Children1 {
			queue = append(queue, NodeWithDepth{node: child, depth: current.depth + 1, parent: current.node.Item1["Name"]})
		}

		// Tambahkan Children2 ke antrian
		for _, child := range current.node.Children2 {
			queue = append(queue, NodeWithDepth{node: child, depth: current.depth + 1, parent: current.node.Item2["Name"]})
		}
	}
}
