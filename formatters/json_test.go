package formatters

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatter(t *testing.T) {
	// --- 1. Создание mock-дерева различий (MapDiff) ---
	// Это должно быть то же дерево, что используется в других тестах,
	// чтобы обеспечить согласованность.
	mockDiff := CreateMockDiff()

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
