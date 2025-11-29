package formatters

import (
	"strings"
	"testing"

	"code/models"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatter(t *testing.T) {
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

	// --- 1. Создание mock-дерева различий (MapDiff) ---
	// Это должно быть то же дерево, что используется в других тестах,
	// чтобы обеспечить согласованность.
	mockDiff := models.MapDiff{
		{
			Key:  "common",
			Type: models.Nested,
			Children: []models.DiffNode{
				{Key: "follow", Type: models.Added, Value2: false},
				{Key: "setting1", Type: models.Unchanged, Value1: "Value 1"},
				{Key: "setting2", Type: models.Removed, Value1: 200.0},
				{Key: "setting3", Type: models.Changed, Value1: true, Value2: nil}, // Changed: true -> null
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

	// --- 2. Ожидаемый результат  ---
	expected := `[
  {
    "Key": "common",
    "Type": "nested",
    "Value1": null,
    "Value2": null,
    "Children": [
      {
        "Key": "follow",
        "Type": "added",
        "Value1": null,
        "Value2": false,
        "Children": null
      },
      {
        "Key": "setting1",
        "Type": "unchanged",
        "Value1": "Value 1",
        "Value2": null,
        "Children": null
      },
      {
        "Key": "setting2",
        "Type": "removed",
        "Value1": 200,
        "Value2": null,
        "Children": null
      },
      {
        "Key": "setting3",
        "Type": "changed",
        "Value1": true,
        "Value2": null,
        "Children": null
      },
      {
        "Key": "setting4",
        "Type": "added",
        "Value1": null,
        "Value2": "blah blah",
        "Children": null
      },
      {
        "Key": "setting5",
        "Type": "added",
        "Value1": null,
        "Value2": {
          "key5": "value5"
        },
        "Children": null
      },
      {
        "Key": "setting6",
        "Type": "nested",
        "Value1": null,
        "Value2": null,
        "Children": [
          {
            "Key": "doge",
            "Type": "nested",
            "Value1": null,
            "Value2": null,
            "Children": [
              {
                "Key": "wow",
                "Type": "changed",
                "Value1": "",
                "Value2": "so much",
                "Children": null
              }
            ]
          },
          {
            "Key": "key",
            "Type": "unchanged",
            "Value1": "value",
            "Value2": null,
            "Children": null
          },
          {
            "Key": "ops",
            "Type": "added",
            "Value1": null,
            "Value2": "vops",
            "Children": null
          }
        ]
      }
    ]
  },
  {
    "Key": "group1",
    "Type": "nested",
    "Value1": null,
    "Value2": null,
    "Children": [
      {
        "Key": "baz",
        "Type": "changed",
        "Value1": "bas",
        "Value2": "bars",
        "Children": null
      },
      {
        "Key": "foo",
        "Type": "unchanged",
        "Value1": "bar",
        "Value2": null,
        "Children": null
      },
      {
        "Key": "nest",
        "Type": "changed",
        "Value1": {
          "key": "value"
        },
        "Value2": "str",
        "Children": null
      }
    ]
  },
  {
    "Key": "group2",
    "Type": "removed",
    "Value1": {
      "abc": 12345,
      "deep": {
        "id": 45
      }
    },
    "Value2": null,
    "Children": null
  },
  {
    "Key": "group3",
    "Type": "added",
    "Value1": null,
    "Value2": {
      "deep": {
        "id": {
          "number": 45
        }
      },
      "fee": 100500
    },
    "Children": null
  }
]`

	// --- 3. Вызов тестируемой функции ---
	actual := JSONFormatter(mockDiff)

	// --- 4. Проверка ---
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "Сгенерированный JSON Diff не соответствует ожидаемому.")
}
