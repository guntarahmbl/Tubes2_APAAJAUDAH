package main

import (
	"log"
	"github.com/guntarahmbl/Tubes2_APAAJAUDAH/BE/Utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Backend(target string, algorithm string, maxRecipe int, data *map[string]interface{}) {
	// 0 : BFS
	// 1 : DFS
	var mode int
	if (algorithm == "bfs"){
		mode = 0
	} else {
		mode = 1;
	}
	start := time.Now()

	recipes, nodeCount, err := utils.GetRecipes(target, mode, maxRecipe)
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}
	time := float64(time.Since(start).Milliseconds())

	// utils.PrintListOfTree(recipes)
	fmt.Printf("Recipes found : %d\n", len(recipes))
	fmt.Printf("Execution time : %f\n", time)
	fmt.Printf("Node Visited : %d\n", nodeCount)

	*data = map[string]interface{}{
		"time":    time,
		"count":   nodeCount,
		"recipes": recipes,
	}
}

func main() {

	// Backend("Alcohol","bfs",1)

	router := gin.Default()

	router.Use(cors.Default()) // mengizinkan semua origin

	router.GET("/api/recipes", func(c *gin.Context) {
		target := c.Query("target")      
		algorithm := c.Query("algorithm")      
		maxRecipe := c.Query("maxRecipe")    
	
		// Konversi cukup yang maxRecipe saja karena dari fe algorithm tuh "bfs" dan "dfs"
		maxRecipeInt, _ := strconv.Atoi(maxRecipe)

		var data map[string]interface{}
		Backend(target,algorithm,maxRecipeInt, &data)
	
		// Kirim isi JSON-nya ke frontend (Rencananya langsung kirim)
		c.JSON(http.StatusOK, data)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	
}
