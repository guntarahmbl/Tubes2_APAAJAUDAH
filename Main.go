package main

import (
	"log"
	"github.com/guntarahmbl/Tubes2_APAAJAUDAH/BE/Utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func Backend(target string, algorithm string, maxRecipe int) {
	// 0 : BFS
	// 1 : DFS
	var mode int
	if (algorithm == "bfs"){
		mode = 0
	} else {
		mode = 1;
	}
	recipes, time, nodeCount, err := utils.GetRecipes(target, mode, maxRecipe)
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	// utils.PrintListOfTree(recipes)
	fmt.Printf("Recipes found : %d\n", len(recipes))
	fmt.Printf("Execution time : %f\n", time)
	fmt.Printf("Node Visited : %d\n", nodeCount)

	savePath := "data/recipes.json"
	utils.SaveRecipes(recipes, time, nodeCount, savePath)
}

func main() {

	// Backend("Sun","bfs",4)

	router := gin.Default()

	router.Use(cors.Default()) // mengizinkan semua origin

	router.GET("/api/recipes", func(c *gin.Context) {
		target := c.Query("target")      
		algorithm := c.Query("algorithm")      
		maxRecipe := c.Query("maxRecipe")    
	
		// Konversi cukup yang maxRecipe saja karena dari fe algorithm tuh "bfs" dan "dfs"
		maxRecipeInt, _ := strconv.Atoi(maxRecipe)

		Backend(target,algorithm,maxRecipeInt)
	
		// Baca isi file recipes.json (Rencananya gausah write file)
		data, err := os.ReadFile("data/recipes.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
			return
		}
	
		// Kirim isi JSON-nya ke frontend (Rencananya langsung kirim)
		c.Data(http.StatusOK, "application/json", data)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	
}
