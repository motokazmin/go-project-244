package code

import (
	"fmt"
	"sort"
)

// DiffNodeType определяет тип изменения узла
type DiffNodeType string

const (
	// Типы узлов:
	Added     DiffNodeType = "added"     // Узел есть только во втором файле
	Removed   DiffNodeType = "removed"   // Узел есть только в первом файле
	Unchanged DiffNodeType = "unchanged" // Узел есть в обоих файлах и значения равны
	Changed   DiffNodeType = "changed"   // Узел есть в обоих файлах, но значения отличаются
	Nested    DiffNodeType = "nested"    // Узел является вложенным объектом (map)
)

// DiffNode представляет один узел в дереве различий
type DiffNode struct {
	Key      string
	Type     DiffNodeType
	Value1   interface{} // Значение из первого файла (используется для Removed, Unchanged, Changed)
	Value2   interface{} // Значение из второго файла (используется для Added, Changed)
	Children []DiffNode  // Используется только для Nested
}

// MapDiff — это итоговое дерево различий (корневой уровень)
type MapDiff []DiffNode

// findUniqueSortedKeys собирает все уникальные ключи из двух карт и сортирует их.
func findUniqueSortedKeys(data1, data2 map[string]interface{}) []string {
	keys := make(map[string]bool)
	if data1 != nil {
		for k := range data1 {
			keys[k] = true
		}
	}
	if data2 != nil {
		for k := range data2 {
			keys[k] = true
		}
	}

	sortedKeys := make([]string, 0, len(keys))
	for k := range keys {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// isMap проверяет, является ли значение картой (объектом).
func isMap(v interface{}) (map[string]interface{}, bool) {
	if m, ok := v.(map[string]interface{}); ok {
		return m, true
	}
	return nil, false
}

// BuildDiff рекурсивно строит дерево различий (MapDiff).
func BuildDiff(data1, data2 map[string]interface{}) MapDiff {
	sortedKeys := findUniqueSortedKeys(data1, data2)
	diff := make(MapDiff, 0, len(sortedKeys))

	for _, key := range sortedKeys {
		val1, exists1 := data1[key]
		val2, exists2 := data2[key]

		node := DiffNode{Key: key}

		// 1. Узел удален (есть только в data1)
		if exists1 && !exists2 {
			node.Type = Removed
			node.Value1 = val1
		} else if !exists1 && exists2 {
			// 2. Узел добавлен (есть только в data2)
			node.Type = Added
			node.Value2 = val2
		} else {
			// 3. Узел есть в обоих файлах
			map1, isMap1 := isMap(val1)
			map2, isMap2 := isMap(val2)

			// 3a. Оба значения являются картами (вложенная структура)
			if isMap1 && isMap2 {
				node.Type = Nested
				node.Children = BuildDiff(map1, map2) // РЕКУРСИЯ!
			} else if fmt.Sprint(val1) == fmt.Sprint(val2) {
				// 3b. Значения равны
				node.Type = Unchanged
				node.Value1 = val1
			} else {
				// 3c. Значения изменились
				node.Type = Changed
				node.Value1 = val1
				node.Value2 = val2
			}
		}

		diff = append(diff, node)
	}

	return diff
}
