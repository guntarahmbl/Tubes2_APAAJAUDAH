package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// fungsi untuk membaca image
func ReadElementsImage(filePath string) (map[string]string, error) {
	// Buka file JSON
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Baca isi file
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parsing JSON ke slice map
	var tempElements []map[string]string
	if err := json.Unmarshal(data, &tempElements); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Konversi ke map[string]string
	elements := make(map[string]string)
	for _, elem := range tempElements {
		name := elem["name"]
		image := elem["image"]
		elements[name] = image
	}

	return elements, nil
}

// fungsi untuk membaca data global resep
func ReadElementsRecipes(filePath string) (map[string][][]string, error) {
	// Buka file JSON
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Baca
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Parse JSON ke dalam map
	var result map[string][][]string
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// fungsi untuk membaca data tier
func ReadElementsTier(filePath string) (map[string]int, error) {
	// Buka file JSON
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Baca
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parsing JSON ke slice map sementara karena gabisa langsung map[string]int
	var tempElements []map[string]interface{}
	if err := json.Unmarshal(data, &tempElements); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Konversi ke map[string]int
	elements := make(map[string]int)
	for _, elem := range tempElements {
		name := elem["nama"].(string)
		tier := int(elem["tier"].(float64)) 
		elements[name] = tier
	}

	return elements, nil
}

// fungsi untuk membaca nama elemen
func ReadElementsName(filePath string) ([]string, error) {
	// Buka file JSON
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Baca isi file
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse JSON ke dalam map
	var result map[string][][]string
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Ambil semua kunci sebagai slice
	var elements []string
	for key := range result {
		elements = append(elements, key)
	}

	return elements, nil
}

// fungsi untuk menyimpan recipes ke file json
func SaveRecipes(trees []*TreeNode, time float64, count int, filename string) error {
	if len(trees) == 0 {
		return nil
	}

	// Bungkus data ke dalam struktur map
	data := map[string]interface{}{
		"time":    time,
		"count":   count,
		"recipes": trees,
	}

	// Serialize ke JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Buat file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Tulis data JSON ke file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

