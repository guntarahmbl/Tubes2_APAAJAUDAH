package utils

import (
	"fmt"
	"log"
	"time"
)


func GenerateRecipesMemoKey(node *TreeNode) string {
	name1 := node.Item1["Name"]
	name2 := node.Item2["Name"]
	return fmt.Sprintf("%s:%s", name1, name2)
}

func GenerateCombinationsMemoKey(node1 *TreeNode, node2 *TreeNode) string {
	name1_1 := node1.Item1["Name"]
	name2_1 := node1.Item2["Name"]
	name1_2 := node2.Item1["Name"]
	name2_2 := node2.Item2["Name"]
	return fmt.Sprintf("%s:%s+%s:%s", name1_1, name2_1, name1_2, name2_2)
}

// menghasilkan seluruh resep
func GenerateRecipesTree(node *TreeNode, countRecipe int) ([]*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	// cek apakah node adalah leaf
	if len(node.Children1) == 0 && len(node.Children2) == 0 {
		return []*TreeNode{node}, 1
	}
	if (len(node.Children1) == 0) {
		node.Children1 = append(node.Children1, NullChild)  // jika tidak memiliki children1, padding dengan nullChild
	}
	if (len(node.Children2) == 0) {
		node.Children2 = append(node.Children2, NullChild)  // jika tidak memiliki children2, padding dengan nullChild
	}

	var combinations, childCombinations1, childCombinations2 []*TreeNode
	var nodeCount1, nodeCount2 int

	count := 1
	nodeCount := 0

	for _, child1 := range node.Children1 {

		key1 := GenerateRecipesMemoKey(child1)
		if r1, found := recipesMemo[key1]; found {
			childCombinations1 = r1
		} else {
			
			childCombinations1, nodeCount1 = GenerateRecipesTree(child1, countRecipe)
			nodeCount += nodeCount1
			recipesMemo[key1] = childCombinations1
		}

		for _, child2 := range node.Children2 {

			keyComb := GenerateCombinationsMemoKey(child1, child2)
	
			if savedComb, found := combinationsMemo[keyComb];found {
				combinations = append(combinations, savedComb...)
				
				break
			} 


			key2 := GenerateRecipesMemoKey(child2)
			if r2, found := recipesMemo[key2]; found {
				childCombinations2 = r2
		

			} else {
		
				childCombinations2, nodeCount2 = GenerateRecipesTree(child2, countRecipe)
				nodeCount += nodeCount2
				recipesMemo[key2] = childCombinations2
		
			}

			var tempCombinations = []*TreeNode{}
			for _, subTree1 := range childCombinations1 {
				for _, subTree2 := range childCombinations2 {

					combinedTree := &TreeNode{
						Item1: map[string]string{"Name": node.Item1["Name"], "Image": node.Item1["Image"]}, 
						Item2: map[string]string{"Name": node.Item2["Name"], "Image": node.Item2["Image"]}, 
						Children1: []*TreeNode{subTree1},
						Children2: []*TreeNode{subTree2},
					}
					tempCombinations = append(tempCombinations, combinedTree)
					combinations = append(combinations, combinedTree)
					count++
					nodeCount++
					if (count > countRecipe) {
						return combinations, nodeCount
					}
				}
			}
			combinationsMemo[keyComb] = tempCombinations
		}
	}
	return combinations, nodeCount
}

// fungsi untuk mengambil semua resep dari suatu elemen
func GetRecipes(name string, method int, maxRecipe int) ([]*TreeNode, float64, int, error) {
	// 0 : BFS
	// 1 : DFS

	// membaca recipe global
	globalRecipes, err := ReadElementsRecipes("data/allElementsRecipes.json")
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	// membaca tier elemen
	tier, err := ReadElementsTier("data/allElementsTiers.json")
	if err != nil {
		log.Fatalf("Failed to read tier: %v", err)
	}

	// membaca image elemen
	img , err := ReadElementsImage("data/allElementsImage.json")
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
	start := time.Now()
	if (method == 0) {
		BuildTreeBFS(globalRecipes, root, tier, img)
	} else if (method == 1) {
		BuildTreeDFS(globalRecipes, root, tier, img)
	} else {
		log.Fatalf("Wrong method!")
	}
	
	recipes, nodeCount := GenerateRecipesTree(root, maxRecipe)
	for _, recipe := range recipes {
		FilterNullChildrens(recipe)  // filter anak yang null
	}
	recipes = FilterAllParents(recipes)
	
	elapsed := time.Since(start).Seconds()
	time := elapsed

	return recipes, time, nodeCount, err
}



