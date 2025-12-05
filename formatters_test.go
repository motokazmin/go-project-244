package code

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// createMockDiff создает mock-дерево различий для использования в тестах
func createMockDiff() mapDiff {
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

	return mapDiff{
		{
			Key:  "common",
			Type: nested,
			Children: []diffNode{
				{Key: "follow", Type: added, Value2: false},
				{Key: "setting1", Type: unchanged, Value1: "Value 1"},
				{Key: "setting2", Type: removed, Value1: 200.0},
				{Key: "setting3", Type: changed, Value1: true, Value2: nil},
				{Key: "setting4", Type: added, Value2: "blah blah"},
				{Key: "setting5", Type: added, Value2: mapSetting5},
				{
					Key:  "setting6",
					Type: nested,
					Children: []diffNode{
						{
							Key:  "doge",
							Type: nested,
							Children: []diffNode{
								{Key: "wow", Type: changed, Value1: "", Value2: "so much"},
							},
						},
						{Key: "key", Type: unchanged, Value1: "value"},
						{Key: "ops", Type: added, Value2: "vops"},
					},
				},
			},
		},
		{
			Key:  "group1",
			Type: nested,
			Children: []diffNode{
				{Key: "baz", Type: changed, Value1: "bas", Value2: "bars"},
				{Key: "foo", Type: unchanged, Value1: "bar"},
				{Key: "nest", Type: changed, Value1: mapGroup1NestOld, Value2: "str"},
			},
		},
		{Key: "group2", Type: removed, Value1: mapGroup2},
		{Key: "group3", Type: added, Value2: mapGroup3},
	}
}

func TestStylishFormatter(t *testing.T) {
	mockDiff := createMockDiff()

	expected := `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`

	actual := formatStylish(mockDiff)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "Сгенерированный Stylish Diff не соответствует ожидаемому.")
}

func TestPlainFormatter(t *testing.T) {
	mockDiff := mapDiff{
		{Key: "follow", Type: added, Value2: false},
		{Key: "host", Type: changed, Value1: "google.com", Value2: "hexlet.io"},
		{Key: "group1", Type: changed, Value1: map[string]interface{}{"key": "val"}, Value2: "str"},
		{Key: "timeout", Type: removed, Value1: 50},
		{Key: "setting3", Type: changed, Value1: true, Value2: nil},

		{Key: "nested", Type: nested, Children: []diffNode{
			{Key: "foo", Type: changed, Value1: "bar", Value2: "changed"},
			{Key: "newKey", Type: added, Value2: "value"},
			{Key: "oldKey", Type: removed, Value1: map[string]interface{}{"x": 1}},
			{Key: "server", Type: unchanged, Value1: "nginx"},
		}},
		{Key: "server", Type: unchanged, Value1: "nginx"},
	}

	expected := `Property 'follow' was added with value: false
Property 'host' was updated. From 'google.com' to 'hexlet.io'
Property 'group1' was updated. From [complex value] to 'str'
Property 'timeout' was removed
Property 'setting3' was updated. From true to null
Property 'nested.foo' was updated. From 'bar' to 'changed'
Property 'nested.newKey' was added with value: 'value'
Property 'nested.oldKey' was removed`

	actual := formatPlain(mockDiff)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual))
}

func TestJSONFormatter(t *testing.T) {
	mockDiff := createMockDiff()

	expected := `{
  "diff": [
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
  ]
}`

	actual := formatJSON(mockDiff)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "Сгенерированный JSON Diff не соответствует ожидаемому.")
}
