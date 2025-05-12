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

func Backend(target string, algorithm int, maxRecipe int) {
	// 0 : BFS
	// 1 : DFS
	recipes, err := utils.GetRecipes(target, algorithm, maxRecipe)
	if err != nil {
		log.Fatalf("Failed to read recipes: %v", err)
	}

	utils.PrintListOfTree(recipes)
	fmt.Printf("Recipes length : %d\n", len(recipes))

	savePath := "data/recipes.json"
	utils.SaveRecipes(recipes, savePath)
}

func main() {
	router := gin.Default()

	router.Use(cors.Default()) // mengizinkan semua origin

	router.GET("/api/recipes", func(c *gin.Context) {
		target := c.Query("target")      
		algorithm := c.Query("algorithm")      
		maxRecipe := c.Query("maxRecipe")    
	
		// konversi string ke int (karena Query selalu string)
		algorithmInt, _ := strconv.Atoi(algorithm)
		maxRecipeInt, _ := strconv.Atoi(maxRecipe)

		Backend(target,algorithmInt,maxRecipeInt)
	
		// Baca isi file recipes.json
		data, err := os.ReadFile("data/recipes.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
			return
		}
	
		// Kirim isi JSON-nya ke frontend
		c.Data(http.StatusOK, "application/json", data)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	
}
