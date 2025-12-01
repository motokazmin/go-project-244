package code

import (
	"reflect"
	"sort"
)

// diffNodeType определяет тип изменения узла
type diffNodeType string

const (
	// Типы узлов:
	added     diffNodeType = "added"     // Узел есть только во втором файле
	removed   diffNodeType = "removed"   // Узел есть только в первом файле
	unchanged diffNodeType = "unchanged" // Узел есть в обоих файлах и значения равны
	changed   diffNodeType = "changed"   // Узел есть в обоих файлах, но значения отличаются
	nested    diffNodeType = "nested"    // Узел является вложенным объектом (map)
)

// diffNode представляет один узел в дереве различий
type diffNode struct {
	Key      string
	Type     diffNodeType
	Value1   interface{} // Значение из первого файла (используется для Removed, Unchanged, Changed)
	Value2   interface{} // Значение из второго файла (используется для Added, Changed)
	Children []diffNode  // Используется только для Nested
}

// mapDiff — это итоговое дерево различий (корневой уровень)
type mapDiff []diffNode

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

