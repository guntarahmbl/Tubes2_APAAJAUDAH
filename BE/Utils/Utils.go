package  utils 

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// menghasilkan seluruh resep
func GenerateAllRecipe(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}

	// cek apakah node adalah leaf
	if len(node.Children1) == 0 && len(node.Children2) == 0 {
		return []*TreeNode{node}
	}
	if (len(node.Children1) == 0) {
		node.Children1 = append(node.Children1, NullChild)  // jika tidak memiliki children1, padding dengan nullChild
	}
	if (len(node.Children2) == 0) {
		node.Children2 = append(node.Children2, NullChild)  // jika tidak memiliki children2, padding dengan nullChild
	}

	var combinations []*TreeNode

	for _, child1 := range node.Children1 {
		for _, child2 := range node.Children2 {

 			childCombinations1 := GenerateAllRecipe(child1)
			childCombinations2 := GenerateAllRecipe(child2)

			for _, subTree1 := range childCombinations1 {
				for _, subTree2 := range childCombinations2 {

					combinedTree := &TreeNode{
						Item1: map[string]string{"Name": node.Item1["Name"], "Image": node.Item1["Image"]}, 
						Item2: map[string]string{"Name": node.Item2["Name"], "Image": node.Item2["Image"]}, 
						Children1: []*TreeNode{subTree1},
						Children2: []*TreeNode{subTree2},
					}
					combinations = append(combinations, combinedTree)
				}
			}
		}
	}

	return combinations
}

// fungsi untuk membaca data global resep
func ReadElementsRecipes(filePath string) (map[string][][]string, error) {
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

// fungsi untuk membaca data tier
func ReadElementsTier(filePath string) (map[string]int, error) {
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

// fungsi untuk membaca data tier
func ReadElementsImage(filePath string) (map[string]int, error) {
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
		name := elem["name"].(string)
		tier := int(elem["image"].(float64)) 
		elements[name] = tier
	}

	return elements, nil
}


// fungsi untuk mengonversi tree ke json
func ConvertTreesToJSON(trees []*TreeNode) (string, error) {
	if trees == nil || len(trees) == 0 {
		return "", nil 
	}

	jsonData, err := json.MarshalIndent(trees, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// fungsi untuk menyimpan tree ke file json
func SaveRecipes(trees []*TreeNode, filename string) error {
	if trees == nil || len(trees) == 0 {
		return nil // 
	}

	jsonData, err := json.MarshalIndent(trees, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// fungsi untuk membersihkan NullChild pada tree
func FilterNullChildrens(node *TreeNode) {
	if node == nil {
		return
	}

	// Filter Children1
	var filteredChildren1 []*TreeNode
	for _, child := range node.Children1 {
		if child.Item1["Name"] != "Null" && child.Item2["Name"] != "Null" {
			filteredChildren1 = append(filteredChildren1, child)
			FilterNullChildrens(child) 
		}
	}
	node.Children1 = filteredChildren1

	// Filter Children2
	var filteredChildren2 []*TreeNode
	for _, child := range node.Children2 {
		if child.Item1["Name"] != "Null" && child.Item2["Name"] != "Null" {
			filteredChildren2 = append(filteredChildren2, child)
			FilterNullChildrens(child) 
		}
	}
	node.Children2 = filteredChildren2
}

func FilterAllParents(nodes []*TreeNode) []*TreeNode {
	var allChildren1 []*TreeNode
	for _, node := range nodes {
		if len(node.Children1) > 0 {
			allChildren1 = append(allChildren1, node.Children1[0])
		}
	}
	return allChildren1
}

// fungsi untuk mengambil semua resep dari suatu elemen
func GetAllRecipes(name string, method int) ([]*TreeNode, error) {
	// 0 : BFS
	// 1 : DFS

	globalRecipes, err := ReadElementsRecipes("data/allElementsRecipes.json")
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	tier, err := ReadElementsTier("data/allElementsTiers.json")
	if err != nil {
		log.Fatalf("Failed to read tier: %v", err)
	}

	// buat root tree
	root := &TreeNode{
			Item1: map[string]string{"Name": name},
			Item2: map[string]string{"Name":  "Null"},
			Children2: []*TreeNode{NullChild},
	}

	// pilih method untuk membangun pohon resep
	if (method == 0) {
		BuildTreeBFS(globalRecipes, root, tier)
	} else if (method == 1) {
		BuildTreeDFS(globalRecipes, root, tier)
	} else {
		log.Fatalf("Wrong method!")
	}

	recipes := GenerateAllRecipe(root)
	for _, recipe := range recipes {
		FilterNullChildrens(recipe)  // filter anak yang null
	}
	recipes = FilterAllParents(recipes)

	return recipes, err
}

// fungsi untuk mengambil sejumlah resep dari suatu elemen
func GetRecipes(name string, method int, maxRecipe int) ([]*TreeNode, error) {

	recipes, err := GetAllRecipes(name, method)
	// fmt.Printf("Recipes length : %d\n", len(recipes))

	maxRecipe = min(maxRecipe, len(recipes)) 
	recipes = recipes[:maxRecipe]  // potong sesuai maxRecipe yang dibutuhkan

	return recipes, err
}
