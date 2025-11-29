package formatters

import (
	"encoding/json"
	"fmt"

	"code/models"
)

// JSONFormatter форматирует дерево различий models.MapDiff
// в виде стандартной JSON-строки.
func JSONFormatter(diff models.MapDiff) string {
	// 1. Кодируем MapDiff в JSON
	data, err := json.MarshalIndent(diff, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling diff to JSON: %v", err)
	}

	return string(data)
}
