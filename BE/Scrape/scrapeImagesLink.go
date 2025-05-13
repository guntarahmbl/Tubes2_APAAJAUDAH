package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"regexp"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type ImageInfo struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func ScrapeImageLinksToJson() {
	urlList := []string{
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Diamond",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Kraken",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Reed",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Water+gun",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var imageLinks []ImageInfo

	for _, url := range urlList {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			images := scrapeImageLinks(url)

			mu.Lock()
			imageLinks = append(imageLinks, images...)
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	// Simpan ke file JSON
	file, err := os.Create(filepath.Join("../../data/images.json"))
	if err != nil {
		fmt.Println("Gagal membuat file JSON:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // agar rapi
	if err := encoder.Encode(imageLinks); err != nil {
		fmt.Println("Gagal encode JSON:", err)
	}
}

func scrapeImageLinks(url string) []ImageInfo {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Gagal mendapatkan URL:", url)
		return nil
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Gagal parse HTML:", err)
		return nil
	}

	var results []ImageInfo

	doc.Find("div.category-page__member-left a").Each(func(i int, s *goquery.Selection) {
		img := s.Find("img")
		if img.Length() > 0 {
			src, _ := img.Attr("src")
			title, exists := s.Attr("title")
			if exists && src != "" && strings.Contains(src, "nocookie.net") {
				cleanSrc := stripImageURL(src)
				results = append(results, ImageInfo{
					Name:  title,
					Image: cleanSrc,
				})
			}
		}
	})

	return results
}

func stripImageURL(url string) string {
	re := regexp.MustCompile(`(https://[^?]+?\.(svg|png|jpg|jpeg))`)
	match := re.FindStringSubmatch(url)
	if len(match) > 0 {
		return match[1]
	}
	return url
}


func main() {
	ScrapeImageLinksToJson()
}