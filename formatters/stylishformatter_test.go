package formatters

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylishFormatter(t *testing.T) {
	// Создание mock-дерева различий (MapDiff) ---
	mockDiff := CreateMockDiff()

	// Ожидаемый результат (Expected Output) ---
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

	// Вызов тестируемой функции ---
	actual := StylishFormatter(mockDiff)

	// Проверка ---
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "Сгенерированный Stylish Diff не соответствует ожидаемому.")
}
