package formatters

import (
	"code/models"
	"fmt"
	"strings"
)

// PlainFormatter форматирует дерево различий в плоском формате.
func PlainFormatter(diff models.MapDiff) string {
	var builder strings.Builder
	// Рекурсивный обход, начиная с пустого пути
	formatNodePlain(diff, "", &builder)

	// Удаляем последний символ новой строки, если он есть
	result := builder.String()
	if strings.HasSuffix(result, "\n") {
		return result[:len(result)-1]
	}
	return result
}

// formatNodePlain рекурсивно обходит дерево и строит вывод в плоском формате.
func formatNodePlain(diff models.MapDiff, path string, builder *strings.Builder) {
	for _, node := range diff {
		currentPath := node.Key
		if path != "" {
			currentPath = path + "." + node.Key
		}

		switch node.Type {
		case models.Added:
			formattedValue := formatPlainValue(node.Value2)
			builder.WriteString(fmt.Sprintf("Property '%s' was added with value: %s\n", currentPath, formattedValue))

		case models.Removed:
			builder.WriteString(fmt.Sprintf("Property '%s' was removed\n", currentPath))

		case models.Changed:
			oldVal := formatPlainValue(node.Value1)
			newVal := formatPlainValue(node.Value2)
			builder.WriteString(fmt.Sprintf("Property '%s' was updated. From %s to %s\n", currentPath, oldVal, newVal))

		case models.Nested:
			// Рекурсивный вызов для вложенных узлов, передавая текущий путь
			formatNodePlain(node.Children, currentPath, builder)

		// Unchanged узлы игнорируются в плоском формате
		case models.Unchanged:
			continue
		}
	}
}

// formatPlainValue форматирует значение для плоского формата.
func formatPlainValue(v interface{}) string {
	if v == nil {
		return "null"
	}
	// Проверяем, является ли значение вложенным объектом
	if _, isMap := models.IsMap(v); isMap {
		return "[complex value]"
	}

	strVal := fmt.Sprint(v)
	// Добавляем кавычки для строковых значений, как требуется по спецификации
	if _, ok := v.(string); ok {
		return fmt.Sprintf("'%s'", strVal)
	}

	return strVal
}
