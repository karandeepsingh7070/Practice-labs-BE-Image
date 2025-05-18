package utils

import (
	"encoding/json"
	"os"

	"github.com/karan/practicelapbs/models"
)

var Problems []models.Problem

func LoadProblemsFromFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &Problems)
}
