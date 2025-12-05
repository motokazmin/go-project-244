package code

import (
	"reflect"
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

// deepEqual проверяет глубокое равенство двух значений
func deepEqual(val1, val2 interface{}) bool {
	return reflect.DeepEqual(val1, val2)
}


