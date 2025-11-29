package code

import (
	"code/models"
	"fmt"
)

// BuildDiff рекурсивно строит дерево различий (MapDiff).
func BuildDiff(data1, data2 map[string]interface{}) models.MapDiff {
	sortedKeys := models.FindUniqueSortedKeys(data1, data2)
	diff := make(models.MapDiff, 0, len(sortedKeys))

	for _, key := range sortedKeys {
		val1, exists1 := data1[key]
		val2, exists2 := data2[key]

		node := models.DiffNode{Key: key}

		// 1. Узел удален (есть только в data1)
		if exists1 && !exists2 {
			node.Type = models.Removed
			node.Value1 = val1
		} else if !exists1 && exists2 {
			// 2. Узел добавлен (есть только в data2)
			node.Type = models.Added
			node.Value2 = val2
		} else {
			// 3. Узел есть в обоих файлах
			map1, isMap1 := models.IsMap(val1)
			map2, isMap2 := models.IsMap(val2)

			// 3a. Оба значения являются картами (вложенная структура)
			if isMap1 && isMap2 {
				node.Type = models.Nested
				node.Children = BuildDiff(map1, map2) // РЕКУРСИЯ!
			} else if fmt.Sprint(val1) == fmt.Sprint(val2) {
				// 3b. Значения равны
				node.Type = models.Unchanged
				node.Value1 = val1
			} else {
				// 3c. Значения изменились
				node.Type = models.Changed
				node.Value1 = val1
				node.Value2 = val2
			}
		}

		diff = append(diff, node)
	}

	return diff
}
