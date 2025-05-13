package utils

import (
	"fmt"
	"strings"
	"sync"
)

// membangun pohon dengan algoritma DFS
func BuildTreeDFS(result map[string][][]string, root *TreeNode, tier map[string]int, img map[string]string) {
    var memoMutex sync.Mutex // Mutex untuk melindungi akses ke memo

    // Fungsi untuk memproses item (item1 atau item2) secara rekursif
    var processItem func(name string, children *[]*TreeNode, tier map[string]int, wg *sync.WaitGroup)
    processItem = func(name string, children *[]*TreeNode, tier map[string]int, wg *sync.WaitGroup) {
        defer wg.Done() // Selesai, kurangi counter WaitGroup

        memoMutex.Lock()
        memoChildren, found := recipesMemo[name]
        memoMutex.Unlock()

        if found {
            *children = append(*children, memoChildren...) // Ambil dari memo jika sudah diproses
            return
        }

        recipes := result[name]
        var newChildren []*TreeNode

        for _, val := range recipes { // Iterasi tiap resep
            
				// jika tier lebih tinggi, skip
				if tier[name] <= tier[val[0]] || tier[name] <= tier[val[1]] {
					continue
				}

				// Jika mengandung time, skip
				if val[0] == "Time" || val[1] == "Time" {
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

            // Rekursi dalam goroutine
            wg.Add(1)
            go processItem(val[0], &newNode.Children1, tier, wg)
            wg.Add(1)
            go processItem(val[1], &newNode.Children2, tier, wg)
        }

        memoMutex.Lock()
        recipesMemo[name] = newChildren // Simpan hasil ke memo
        memoMutex.Unlock()

        *children = append(*children, newChildren...)
    }

    // Proses Item1 dan Item2 secara paralel
    var wg sync.WaitGroup

    wg.Add(1)
    go processItem(root.Item1["Name"], &root.Children1, tier, &wg)

    wg.Add(1)
    go processItem(root.Item2["Name"], &root.Children2, tier, &wg)

    wg.Wait() // Tunggu semua goroutine selesai
}

// melakukan traverse dengan DFS
func TraverseTreeDFS(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	// Cetak informasi node saat ini
	fmt.Printf("%s %s, %s \n", strings.Repeat("-", depth), node.Item1, node.Item2)

	// Rekursi untuk Children1
	for _, child := range node.Children1 {
		TraverseTreeDFS(child, depth+1)
	}

	// Rekursi untuk Children2
	for _, child := range node.Children2 {
		TraverseTreeDFS(child, depth+1)
	}
}

