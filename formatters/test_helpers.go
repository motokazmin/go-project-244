package formatters

import (
	"code/models"
)

// CreateMockDiff создает mock-дерево различий для использования в тестах
func CreateMockDiff() models.MapDiff {
	// Маппы для имитации комплексных вложенных значений (float64, как после парсинга JSON)
	mapSetting5 := map[string]interface{}{"key5": "value5"}
	mapGroup2 := map[string]interface{}{
		"abc": 12345.0,
		"deep": map[string]interface{}{
			"id": 45.0,
		},
	}
	mapGroup3 := map[string]interface{}{
		"deep": map[string]interface{}{
			"id": map[string]interface{}{
				"number": 45.0,
			},
		},
		"fee": 100500.0,
	}
	mapGroup1NestOld := map[string]interface{}{"key": "value"}

	return models.MapDiff{
		{
			Key:  "common",
			Type: models.Nested,
			Children: []models.DiffNode{
				{Key: "follow", Type: models.Added, Value2: false},
				{Key: "setting1", Type: models.Unchanged, Value1: "Value 1"},
				{Key: "setting2", Type: models.Removed, Value1: 200.0},
				{Key: "setting3", Type: models.Changed, Value1: true, Value2: nil},
				{Key: "setting4", Type: models.Added, Value2: "blah blah"},
				{Key: "setting5", Type: models.Added, Value2: mapSetting5},
				{
					Key:  "setting6",
					Type: models.Nested,
					Children: []models.DiffNode{
						{
							Key:  "doge",
							Type: models.Nested,
							Children: []models.DiffNode{
								{Key: "wow", Type: models.Changed, Value1: "", Value2: "so much"},
							},
						},
						{Key: "key", Type: models.Unchanged, Value1: "value"},
						{Key: "ops", Type: models.Added, Value2: "vops"},
					},
				},
			},
		},
		{
			Key:  "group1",
			Type: models.Nested,
			Children: []models.DiffNode{
				{Key: "baz", Type: models.Changed, Value1: "bas", Value2: "bars"},
				{Key: "foo", Type: models.Unchanged, Value1: "bar"},
				{Key: "nest", Type: models.Changed, Value1: mapGroup1NestOld, Value2: "str"},
			},
		},
		{Key: "group2", Type: models.Removed, Value1: mapGroup2},
		{Key: "group3", Type: models.Added, Value2: mapGroup3},
	}
}

