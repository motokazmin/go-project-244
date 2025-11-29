package formatters

import (
	"code/models"
	"fmt"
	"strings"
)

// StylishFormatter форматирует дерево различий в стиле Stylish.
func StylishFormatter(diff models.MapDiff) string {
	var builder strings.Builder
	builder.WriteString("{\n")
	formatNodes(diff, &builder, 1) // Начинаем с отступа 1 (4 пробела)
	builder.WriteString("}")
	return builder.String()
}

// formatNodes рекурсивно обходит узлы и записывает отформатированную строку.
// indentLevel - уровень вложенности для отступов (1 уровень = 4 пробела)
func formatNodes(diff models.MapDiff, builder *strings.Builder, indentLevel int) {
	const indentSize = 4

	// Базовый отступ: (4 * уровень) - 2 пробела (для символов "+", "-", " ")
	baseIndent := strings.Repeat(" ", indentLevel*indentSize-2)

	for _, node := range diff {
		switch node.Type {
		case models.Unchanged:
			// Отступ + "  " + ключ: значение
			builder.WriteString(fmt.Sprintf("%s  %s: %v\n", baseIndent, node.Key, formatValue(node.Value1, indentLevel)))
		case models.Added:
			// Отступ + "+ " + ключ: значение
			builder.WriteString(fmt.Sprintf("%s+ %s: %v\n", baseIndent, node.Key, formatValue(node.Value2, indentLevel)))
		case models.Removed:
			// Отступ + "- " + ключ: значение
			builder.WriteString(fmt.Sprintf("%s- %s: %v\n", baseIndent, node.Key, formatValue(node.Value1, indentLevel)))
		case models.Changed:
			// Сначала Removed, потом Added
			builder.WriteString(fmt.Sprintf("%s- %s: %v\n", baseIndent, node.Key, formatValue(node.Value1, indentLevel)))
			builder.WriteString(fmt.Sprintf("%s+ %s: %v\n", baseIndent, node.Key, formatValue(node.Value2, indentLevel)))
		case models.Nested:
			// Вложенная структура: открываем новую секцию
			builder.WriteString(fmt.Sprintf("%s  %s: {\n", baseIndent, node.Key))
			// Рекурсивный вызов для дочерних узлов
			formatNodes(node.Children, builder, indentLevel+1)
			// Закрывающая скобка с правильным отступом
			builder.WriteString(fmt.Sprintf("%s  }\n", baseIndent))
		}
	}
}

// formatValue форматирует значение. Если значение - карта (объект),
// он рекурсивно выводит его содержимое в стиле Stylish.
func formatValue(v interface{}, indentLevel int) string {
	if v == nil {
		return "null"
	}

	if m, isMap := models.IsMap(v); isMap {
		var innerBuilder strings.Builder
		innerBuilder.WriteString("{\n")

		// Отступы для содержимого вложенной карты
		nextBaseIndent := strings.Repeat(" ", (indentLevel+1)*4-2)

		// Получаем и сортируем ключи для детерминированного вывода
		keys := models.FindUniqueSortedKeys(m, nil)

		for _, key := range keys {
			// Рекурсивно форматируем вложенные значения (с увеличенным уровнем отступа)
			innerValue := formatValue(m[key], indentLevel+1)
			// Внутри вложенного объекта все узлы считаются Unchanged (нет префиксов "+/-")
			innerBuilder.WriteString(fmt.Sprintf("%s  %s: %v\n", nextBaseIndent, key, innerValue))
		}

		// Закрывающая скобка с отступом текущего уровня + 2 пробела для "  "
		closingIndent := strings.Repeat(" ", indentLevel*4)
		innerBuilder.WriteString(fmt.Sprintf("%s}", closingIndent))

		return innerBuilder.String()
	}
	// Для простых значений используем стандартное строковое представление
	return fmt.Sprint(v)
}
