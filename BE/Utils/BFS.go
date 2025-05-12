package utils

import (
	"fmt"
	"strings"
	"sync"
)

// membangun pohon dengan algoritma BFS
func BuildTreeBFS(result map[string][][]string, root *TreeNode, tier map[string]int, img map[string]string) {
	// Queue untuk BFS
	queue := []*TreeNode{root}

	// Mutex untuk memproses memo secara thread-safe
	var memoMutex sync.Mutex
	memo := make(map[string][]*TreeNode)

	// Proses BFS
	for len(queue) > 0 {
		currentNode := queue[0] // Ambil node pertama dalam queue
		queue = queue[1:]      // Dequeue

		var wg sync.WaitGroup // WaitGroup untuk menunggu semua goroutine selesai

		// Fungsi untuk memproses item (item1 atau item2)
		processItem := func(name string, children *[]*TreeNode, currentQueue *[]*TreeNode) {
			defer wg.Done() // Defer untuk mengurangi WaitGroup counter
			memoMutex.Lock()
			memoChildren, found := memo[name]
			memoMutex.Unlock()

			if found {
				*children = append(*children, memoChildren...) // Ambil dari memo jika sudah diproses
				return
			}

			recipes := result[name]
			var newChildren []*TreeNode

			for _, val := range recipes {
				if tier[name] <= tier[val[0]] || tier[name] <= tier[val[1]] {
					continue
				}

				newNode := &TreeNode{
					Item1: map[string]string{
						"Name":  val[0],
						"Image": img[val[0]],
					},
					Item2: map[string]string{
						"Name":  val[1],
						"Image": img[val[1]],
					},
				}

				newChildren = append(newChildren, newNode)
				*currentQueue = append(*currentQueue, newNode) // Tambahkan ke queue
			}

			memoMutex.Lock()
			memo[name] = newChildren // Simpan hasil ke memo
			memoMutex.Unlock()

			*children = append(*children, newChildren...)
		}

		// Proses item1 secara paralel
		wg.Add(1)
		go processItem(currentNode.Item1["Name"], &currentNode.Children1, &queue)

		// Proses item2 secara paralel
		wg.Add(1)
		go processItem(currentNode.Item2["Name"], &currentNode.Children2, &queue)

		wg.Wait() // Tunggu semua goroutine selesai
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
