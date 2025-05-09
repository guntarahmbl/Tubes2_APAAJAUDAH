package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Nama   string `json:"nama"`
	Gambar string `json:"gambar"`
}

func ScrapeElementNames() {
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)?action=edit"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Gagal mengambil halaman:", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Gagal membaca dokumen HTML:", err)
	}

	text := doc.Find("textarea").Text()

	// Regex untuk menangkap semua pola {{Icon2|NamaElemen}}
	re := regexp.MustCompile(`\{\{Icon2\|([^\}]+)\}\}`)
	matches := re.FindAllStringSubmatch(text, -1)

	elementMap := make(map[string]bool)
	for _, match := range matches {
		element := strings.ToLower(strings.TrimSpace(match[1]))
		elementMap[element] = true
	}

	// Buat list Element
	var elements []Element
	for element := range elementMap {
		imgName := strings.ReplaceAll(element, " ", "_")
		imgName = strings.Title(imgName) 
		imgFileName := imgName + ".png" 

		name := strings.Title(element)
		e := Element{
			Nama:   name,
			Gambar: "/images/" + imgFileName,
		}
		elements = append(elements, e)
	}

	// Urutkan berdasarkan nama
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Nama < elements[j].Nama
	})

	// Simpan ke file JSON
	file, err := os.Create("allElements.json")
	if err != nil {
		log.Fatal("Gagal membuat file JSON:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(elements); err != nil {
		log.Fatal("Gagal menyimpan data JSON:", err)
	}

	log.Println("Sukses menyimpan data ke allElements.json")
}
