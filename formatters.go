package code

import (
	"encoding/json"
	"fmt"
	"strings"
)

// formatStylish форматирует дерево различий в стиле Stylish.
func formatStylish(diff mapDiff) string {
	var builder strings.Builder
	builder.WriteString("{\n")
	formatStylishNodes(diff, &builder, 1)
	builder.WriteString("}")
	return builder.String()
}

// formatStylishNodes рекурсивно обходит узлы и записывает отформатированную строку.
// indentLevel - уровень вложенности для отступов (1 уровень = 4 пробела)
func formatStylishNodes(diff mapDiff, builder *strings.Builder, indentLevel int) {
	const indentSize = 4

	// Базовый отступ: (4 * уровень) - 2 пробела (для символов "+", "-", " ")
	baseIndent := strings.Repeat(" ", indentLevel*indentSize-2)

	for _, node := range diff {
		switch node.Type {
		case unchanged:
			// Отступ + "  " + ключ: значение
			fmt.Fprintf(builder, "%s  %s: %v\n", baseIndent, node.Key, formatStylishValue(node.Value1, indentLevel))
		case added:
			// Отступ + "+ " + ключ: значение
			fmt.Fprintf(builder, "%s+ %s: %v\n", baseIndent, node.Key, formatStylishValue(node.Value2, indentLevel))
		case removed:
			// Отступ + "- " + ключ: значение
			fmt.Fprintf(builder, "%s- %s: %v\n", baseIndent, node.Key, formatStylishValue(node.Value1, indentLevel))
		case changed:
			// Сначала Removed, потом Added
			fmt.Fprintf(builder, "%s- %s: %v\n", baseIndent, node.Key, formatStylishValue(node.Value1, indentLevel))
			fmt.Fprintf(builder, "%s+ %s: %v\n", baseIndent, node.Key, formatStylishValue(node.Value2, indentLevel))
		case nested:
			// Вложенная структура: открываем новую секцию
			fmt.Fprintf(builder, "%s  %s: {\n", baseIndent, node.Key)
			// Рекурсивный вызов для дочерних узлов
			formatStylishNodes(node.Children, builder, indentLevel+1)
			// Закрывающая скобка с правильным отступом
			fmt.Fprintf(builder, "%s  }\n", baseIndent)
		}
	}
}

// formatStylishValue форматирует значение. Если значение - карта (объект),
// он рекурсивно выводит его содержимое в стиле Stylish.
func formatStylishValue(v interface{}, indentLevel int) string {
	if v == nil {
		return "null"
	}

	if m, isMap := isMap(v); isMap {
		var innerBuilder strings.Builder
		innerBuilder.WriteString("{\n")

		// Отступы для содержимого вложенной карты
		nextBaseIndent := strings.Repeat(" ", (indentLevel+1)*4-2)

		// Получаем и сортируем ключи для детерминированного вывода
		keys := findUniqueSortedKeys(m, nil)

		for _, key := range keys {
			// Рекурсивно форматируем вложенные значения (с увеличенным уровнем отступа)
			innerValue := formatStylishValue(m[key], indentLevel+1)
			// Внутри вложенного объекта все узлы считаются Unchanged (нет префиксов "+/-")
			fmt.Fprintf(&innerBuilder, "%s  %s: %v\n", nextBaseIndent, key, innerValue)
		}

		// Закрывающая скобка с отступом текущего уровня + 2 пробела для "  "
		closingIndent := strings.Repeat(" ", indentLevel*4)
		fmt.Fprintf(&innerBuilder, "%s}", closingIndent)

		return innerBuilder.String()
	}
	// Для простых значений используем стандартное строковое представление
	return fmt.Sprint(v)
}

// formatPlain форматирует дерево различий в плоском формате.
func formatPlain(diff mapDiff) string {
	var builder strings.Builder
	// Рекурсивный обход, начиная с пустого пути
	formatPlainNodes(diff, "", &builder)

	// Удаляем последний символ новой строки, если он есть
	result := builder.String()
	if strings.HasSuffix(result, "\n") {
		return result[:len(result)-1]
	}
	return result
}

// formatPlainNodes рекурсивно обходит дерево и строит вывод в плоском формате.
func formatPlainNodes(diff mapDiff, path string, builder *strings.Builder) {
	for _, node := range diff {
		currentPath := node.Key
		if path != "" {
			currentPath = path + "." + node.Key
		}

		switch node.Type {
		case added:
			formattedValue := formatPlainValue(node.Value2)
			fmt.Fprintf(builder, "Property '%s' was added with value: %s\n", currentPath, formattedValue)

		case removed:
			fmt.Fprintf(builder, "Property '%s' was removed\n", currentPath)

		case changed:
			oldVal := formatPlainValue(node.Value1)
			newVal := formatPlainValue(node.Value2)
			fmt.Fprintf(builder, "Property '%s' was updated. From %s to %s\n", currentPath, oldVal, newVal)

		case nested:
			// Рекурсивный вызов для вложенных узлов, передавая текущий путь
			formatPlainNodes(node.Children, currentPath, builder)

		// Unchanged узлы игнорируются в плоском формате
		case unchanged:
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
	if _, isMapVal := isMap(v); isMapVal {
		return "[complex value]"
	}

	strVal := fmt.Sprint(v)
	// Добавляем кавычки для строковых значений, как требуется по спецификации
	if _, ok := v.(string); ok {
		return fmt.Sprintf("'%s'", strVal)
	}

	return strVal
}

// formatJSON форматирует дерево различий в JSON формате.
func formatJSON(diff mapDiff) string {
	// Оборачиваем mapDiff в объект для корректного JSON структуры
	// Вместо массива [ ... ] возвращаем объект { "diff": [ ... ] }
	result := map[string]interface{}{
		"diff": diff,
	}
	
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling diff to JSON: %v", err)
	}

	return string(data)
}
