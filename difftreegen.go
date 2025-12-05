package code

// buildDiff рекурсивно строит дерево различий (MapDiff).
func buildDiff(data1, data2 map[string]interface{}) MapDiff {
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
				node.Children = buildDiff(map1, map2) // РЕКУРСИЯ!
			} else if deepEqual(val1, val2) {
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
