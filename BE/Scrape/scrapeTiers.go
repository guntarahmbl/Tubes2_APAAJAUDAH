package scrape

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeElementTiers mengembalikan map yang berisi nama elemen sebagai key dan tier sebagai nilai.
func ScrapeElementTiers() (map[string]int, error) {
	// URL yang akan di-scrape
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)?action=edit"

	// Mengambil HTML dari URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching the page: %w", err)
	}
	defer resp.Body.Close()

	// Menggunakan goquery untuk parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the HTML: %w", err)
	}

	// Map untuk menampung elemen dan tier (key: nama elemen, value: tier)
	elements := make(map[string]int)

	// Regular expression untuk menemukan elemen dan tier
	re := regexp.MustCompile(`=== Tier (\d+) elements ===`)

	// Parse setiap baris tabel
	doc.Find(".list-table").Each(func(i int, s *goquery.Selection) {
		// Dapatkan tier
		tierText := re.FindStringSubmatch(s.Text())
		if len(tierText) > 1 {
			tier := tierText[1] // Ambil tier dari regex

			// Konversi tier ke integer
			tierInt := 0
			fmt.Sscanf(tier, "%d", &tierInt)

			// Parse setiap baris dalam tabel
			s.Find("tr").Each(func(i int, row *goquery.Selection) {
				cells := row.Find("td")
				if cells.Length() > 0 {
					// Ambil nama elemen dari kolom pertama (Icon2)
					name := cells.First().Text()
					// Menghapus spasi ekstra yang tidak diperlukan
					name = strings.TrimSpace(name)

					// Menambahkan elemen dan tier ke map
					elements[name] = tierInt
				}
			})
		}
	})

	return elements, nil
}