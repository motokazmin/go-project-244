package formatters

import (
	"code/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainFormatter(t *testing.T) {
	mockDiff := models.MapDiff{
		{Key: "follow", Type: models.Added, Value2: false},
		{Key: "host", Type: models.Changed, Value1: "google.com", Value2: "hexlet.io"},
		{Key: "group1", Type: models.Changed, Value1: map[string]interface{}{"key": "val"}, Value2: "str"},
		{Key: "timeout", Type: models.Removed, Value1: 50},
		{Key: "setting3", Type: models.Changed, Value1: true, Value2: nil}, // Проверка nil -> null

		{Key: "nested", Type: models.Nested, Children: []models.DiffNode{
			{Key: "foo", Type: models.Changed, Value1: "bar", Value2: "changed"},
			{Key: "newKey", Type: models.Added, Value2: "value"},
			{Key: "oldKey", Type: models.Removed, Value1: map[string]interface{}{"x": 1}},
			{Key: "server", Type: models.Unchanged, Value1: "nginx"},
		}},
		// Узел неизменен (игнорируется)
		{Key: "server", Type: models.Unchanged, Value1: "nginx"},
	}

	// Ожидаемый результат в Плоском формате
	expected := `Property 'follow' was added with value: false
Property 'host' was updated. From 'google.com' to 'hexlet.io'
Property 'group1' was updated. From [complex value] to 'str'
Property 'timeout' was removed
Property 'setting3' was updated. From true to null
Property 'nested.foo' was updated. From 'bar' to 'changed'
Property 'nested.newKey' was added with value: 'value'
Property 'nested.oldKey' was removed`

	actual := PlainFormatter(mockDiff)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual))
}
