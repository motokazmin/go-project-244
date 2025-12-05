package code

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
		log.Printf("[DEBUG] parseJSON: первая попытка (map) провалилась: %v", err)
		// Если это не объект, пробуем распарсить как массив
		var arrayResult []interface{}
		arrayErr := json.Unmarshal(data, &arrayResult)
		if arrayErr == nil {
			log.Printf("[DEBUG] parseJSON: успешно распарсили как массив, возвращаем пустую map")
			// Успешно распарсили массив, возвращаем пустую map
			return make(map[string]interface{}), nil
		}
		log.Printf("[DEBUG] parseJSON: вторая попытка (array) тоже провалилась: %v", arrayErr)
		// Если оба варианта не сработали, возвращаем оригинальную ошибку
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	// Проверяем, что result не nil (был JSON null)
	if result == nil {
		log.Printf("[DEBUG] parseJSON: result == nil (был null)")
		return nil, fmt.Errorf("invalid JSON: null is not allowed")
	}

	log.Printf("[DEBUG] parseJSON: успешно распарсили как объект")
	return result, nil
}

// parseYAML парсит YAML данные
func parseYAML(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		log.Printf("[DEBUG] parseYAML: первая попытка (map) провалилась: %v", err)
		// Если это не объект, пробуем распарсить как массив
		var arrayResult []interface{}
		arrayErr := yaml.Unmarshal(data, &arrayResult)
		if arrayErr == nil {
			log.Printf("[DEBUG] parseYAML: успешно распарсили как массив, возвращаем пустую map")
			// Успешно распарсили массив, возвращаем пустую map
			return make(map[string]interface{}), nil
		}
		log.Printf("[DEBUG] parseYAML: вторая попытка (array) тоже провалилась: %v", arrayErr)
		// Если оба варианта не сработали, возвращаем оригинальную ошибку
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}

	// Проверяем, что result не nil (был YAML null)
	if result == nil {
		log.Printf("[DEBUG] parseYAML: result == nil (был null)")
		return nil, fmt.Errorf("invalid YAML: null is not allowed")
	}

	log.Printf("[DEBUG] parseYAML: успешно распарсили как объект")
	return result, nil
}

// GenDiff — главная функция, которая собирает все части
func GenDiff(path1, path2, formatter string) (string, error) {
	log.Printf("[DEBUG] GenDiff called: path1=%s, path2=%s, formatter=%s", path1, path2, formatter)
	
	data1, err := parseFile(path1)
	if err != nil {
		log.Printf("[DEBUG] GenDiff: ошибка при parseFile(path1): %v", err)
		return "", err
	}
	log.Printf("[DEBUG] GenDiff: успешно распарсили path1")

	data2, err := parseFile(path2)
	if err != nil {
		log.Printf("[DEBUG] GenDiff: ошибка при parseFile(path2): %v", err)
		return "", err
	}
	log.Printf("[DEBUG] GenDiff: успешно распарсили path2")

	diffTree := buildDiff(data1, data2)
	log.Printf("[DEBUG] GenDiff: успешно построили diffTree")

	switch formatter {
	case "stylish":
		log.Printf("[DEBUG] GenDiff: используем stylish formatter")
		return formatStylish(diffTree), nil
	case "plain":
		log.Printf("[DEBUG] GenDiff: используем plain formatter")
		return formatPlain(diffTree), nil
	case "json":
		log.Printf("[DEBUG] GenDiff: используем json formatter")
		return formatJSON(diffTree), nil
	default:
		log.Printf("[DEBUG] GenDiff: неподдерживаемый formatter: %s", formatter)
		return "", fmt.Errorf("unsupported formatter %s", formatter)
	}
}
