package utils

import (
	"encoding/json"
)

// fungsi untuk mengonversi tree ke json
func ConvertTreesToJSON(trees []*TreeNode) (string, error) {
	if len(trees) == 0 {
		return "", nil 
	}

	jsonData, err := json.MarshalIndent(trees, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}