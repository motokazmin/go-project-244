package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseFile читает и парсит файл в map[string]interface{}
// Формат определяется по расширению файла
func ParseFile(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".json":
		return parseJSON(data)
	case ".yaml", ".yml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %s (supported: .json, .yaml, .yml)", ext)
	}
}

// parseJSON парсит JSON данные
func parseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return result, nil
}

// parseYAML парсит YAML данные
func parseYAML(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	return result, nil
}

// GenDiff генерирует строку с различиями между двумя map
func GenDiff(data1, data2 map[string]interface{}) string {
	keys := make(map[string]bool)
	for k := range data1 {
		keys[k] = true
	}
	for k := range data2 {
		keys[k] = true
	}

	sortedKeys := make([]string, 0, len(keys))
	for k := range keys {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	var result strings.Builder
	result.WriteString("{\n")

	for _, key := range sortedKeys {
		val1, exists1 := data1[key]
		val2, exists2 := data2[key]

		if exists1 && exists2 {
			if fmt.Sprint(val1) == fmt.Sprint(val2) {
				result.WriteString(fmt.Sprintf("    %s: %v\n", key, val1))
			} else {
				result.WriteString(fmt.Sprintf("  - %s: %v\n", key, val1))
				result.WriteString(fmt.Sprintf("  + %s: %v\n", key, val2))
			}
		} else if exists1 {
			result.WriteString(fmt.Sprintf("  - %s: %v\n", key, val1))
		} else {
			result.WriteString(fmt.Sprintf("  + %s: %v\n", key, val2))
		}
	}

	result.WriteString("}")
	return result.String()
}
