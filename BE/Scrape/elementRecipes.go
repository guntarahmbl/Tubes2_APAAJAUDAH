package scrape

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeRecipes() {
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)?action=edit"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Gagal mengambil halaman:", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca dokumen HTML:", err)
	}

	// Ambil isi textarea yang berisi source wikinya
	text := doc.Find("textarea").Text()

	// Regex untuk mencari data elemen dan resep
	re := regexp.MustCompile(`\{\{Icon2\|([^\}]+)\}\}.*?\{\{RecipesT2\|([^\}]+)\}\}`)

	elementRecipes := make(map[string][][]string)

	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		element := match[1]
		rawRecipes := strings.Split(match[2], "|")
		var recipes [][]string
		for i := 0; i < len(rawRecipes)-1; i += 2 {
			first := strings.Title(strings.TrimSpace(rawRecipes[i]))
			second := strings.Title(strings.TrimSpace(rawRecipes[i+1]))
			if first != "" && second != "" {
				recipes = append(recipes, []string{first, second})
			}
		}
		if len(recipes) > 0 {
			elementRecipes[element] = recipes
		}
	}

	// Simpan ke JSON
	file, err := os.Create("elements_recipes.json")
	if err != nil {
		fmt.Println("Gagal membuat file JSON:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(elementRecipes); err != nil {
		log.Fatal("Gagal menyimpan ke file JSON:", err)
	}

	fmt.Println("Data disimpan ke 'elements_recipes.json'")
}


