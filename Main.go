package main

import (
	"log"
	"github.com/guntarahmbl/Tubes2_APAAJAUDAH/BE/Utils"
	"fmt"
)

func main() {
	// 0 : BFS
	// 1 : DFS
	recipes, err := utils.GetRecipes("Human", 0, 10)
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	utils.PrintListOfTree(recipes)
	fmt.Printf("Recipes length : %d\n", len(recipes))

	savePath := "data/recipes.json"
	utils.SaveRecipes(recipes, savePath)
}
