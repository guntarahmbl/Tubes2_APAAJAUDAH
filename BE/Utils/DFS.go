package utils

import (
	"fmt"
	"strings"
)

// membangun pohon dengan algoritma DFS
func BuildTreeDFS(result map[string][][]string, root *TreeNode, tier map[string]int) {
    // Proses Item1
    if memoChildren, found := memo[root.Item1["Name"]]; !found {
        recipes1 := result[root.Item1["Name"]] // Cari resep-resep Item1

        for _, val := range recipes1 { // Iterasi tiap resep
            // Jika tier sama atau lebih rendah, abaikan
            if tier[root.Item1["Name"]] <= tier[val[0]] || tier[root.Item1["Name"]] <= tier[val[1]] {
                continue
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
            root.Children1 = append(root.Children1, newNode) // Tambahkan sebagai Children1

            // Rekursi untuk node baru
            BuildTreeDFS(result, newNode, tier)
        }

        memo[root.Item1["Name"]] = root.Children1 // Simpan ke memo setelah selesai diproses
    } else {
        root.Children1 = append(root.Children1, memoChildren...) // Ambil dari memo jika sudah pernah diproses
    }

    // Proses Item2
    if memoChildren, found := memo[root.Item2["Name"]]; !found {
        recipes2 := result[root.Item2["Name"]] // Cari resep-resep Item2

        for _, val := range recipes2 { // Iterasi tiap resep
            // Jika tier sama atau lebih rendah, abaikan
            if tier[root.Item2["Name"]] <= tier[val[0]] || tier[root.Item2["Name"]] <= tier[val[1]] {
                continue
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
            root.Children2 = append(root.Children2, newNode) // Tambahkan sebagai Children2

            // Rekursi untuk node baru
            BuildTreeDFS(result, newNode, tier)
        }

        memo[root.Item2["Name"]] = root.Children2 // Simpan ke memo setelah selesai diproses
    } else {
        root.Children2 = append(root.Children2, memoChildren...) // Ambil dari memo jika sudah pernah diproses
    }
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

