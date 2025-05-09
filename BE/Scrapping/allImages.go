package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    // "strings"
	"sync"
    "github.com/PuerkitoBio/goquery"
)

func scrapeImages() {
    urlList := []string{
        "https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2",
        "https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Diamond",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Kraken",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Reed",
		"https://little-alchemy.fandom.com/wiki/Category:Little_Alchemy_2?from=Water+gun",
    }

    os.MkdirAll("images", os.ModePerm)

    var wg sync.WaitGroup

    for _, url := range urlList {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            scrapePage(url)
        }(url)
    }

    wg.Wait() 
}

func scrapePage(url string) {
    res, err := http.Get(url)
    if err != nil {
        fmt.Println("Gagal mendapatkan url:", url)
        return
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        fmt.Println("Gagal parse html:", err)
        return
    }

    // Cari <a title="..."><img src=...></a>
    doc.Find("div.category-page__member-left a").Each(func(i int, s *goquery.Selection) {
        img := s.Find("img")
        if img.Length() > 0 {
            src, _ := img.Attr("src")
            title, exists := s.Attr("title")
            if exists && src != "" {
                filename := filepath.Join("images", title +".png")
                fmt.Printf("Downloading %s -> %s\n", title, filename)
                downloadImage(src, filename)
            }
        }
    })
}

// download gambar
func downloadImage(url, filepath string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(" Gagal download:", url)
        return
    }
    defer resp.Body.Close()

    file, err := os.Create(filepath)
    if err != nil {
        fmt.Println(" Gagal membuat file:", filepath)
        return
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Println(" Gagal menyimpan image:", filepath)
        return
    }

    fmt.Println(" Saved:", filepath)
}
