package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	f "code/formatters"

	"gopkg.in/yaml.v3"
)

// parseFile читает и парсит файл в map[string]interface{}
// Формат определяется по расширению файла
func parseFile(path string) (map[string]interface{}, error) {
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
		// Если это не объект, пробуем распарсить как массив
		var arrayResult []interface{}
		if arrayErr := json.Unmarshal(data, &arrayResult); arrayErr == nil {
			// Успешно распарсили массив, оборачиваем его в map
			return map[string]interface{}{
				"_array": arrayResult,
			}, nil
		}
		// Если оба варианта не сработали, возвращаем оригинальную ошибку
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return result, nil
}

// parseYAML парсит YAML данные
func parseYAML(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		// Если это не объект, пробуем распарсить как массив
		var arrayResult []interface{}
		if arrayErr := yaml.Unmarshal(data, &arrayResult); arrayErr == nil {
			// Успешно распарсили массив, оборачиваем его в map
			return map[string]interface{}{
				"_array": arrayResult,
			}, nil
		}
		// Если оба варианта не сработали, возвращаем оригинальную ошибку
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	return result, nil
}

// GenDiff — главная функция, которая собирает все части
func GenDiff(path1, path2, formater string) (string, error) {
	data1, err := parseFile(path1)
	if err != nil {
		return "", err
	}

	data2, err := parseFile(path2)
	if err != nil {
		return "", err
	}

	diffTree := BuildDiff(data1, data2)

	if formater == "stylish" {
		return f.StylishFormatter(diffTree), nil
	}
	if formater == "plain" {
		return f.PlainFormatter(diffTree), nil
	}
	if formater == "json" {
		return f.JSONFormatter(diffTree), nil
	}

	return "", fmt.Errorf("unsupported formater %s", formater)
}
