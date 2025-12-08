package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const JSON1 = "file1.json"
const JSON2 = "file2.json"
const YAML1 = "file1.yaml"
const YAML2 = "file2.yaml"
const ERRMSG = "GenDiff должна была успешно выполниться"

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name      string
		file1Path string
		file2Path string
		expected  string
	}{
		{
			name:      "main example",
			file1Path: filepath.Join("testdata", JSON1),
			file2Path: filepath.Join("testdata", JSON2),
			expected: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenDiff(tt.file1Path, tt.file2Path, "stylish")
			require.NoError(t, err, ERRMSG)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenDiffEqualFiles(t *testing.T) {
	tests := []struct {
		name      string
		file1Path string
		file2Path string
		expected  string
	}{
		{
			name:      "main example",
			file1Path: filepath.Join("testdata", JSON1),
			file2Path: filepath.Join("testdata", JSON1),
			expected: `{
    follow: false
    host: hexlet.io
    proxy: 123.234.53.22
    timeout: 50
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenDiff(tt.file1Path, tt.file2Path, "stylish")
			require.NoError(t, err, ERRMSG)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Интеграционные тесты
func TestGenDiffIntegration(t *testing.T) {
	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	// Определяем тестовые случаи (Table-Driven Tests)
	tests := []struct {
		name      string
		file1Name string
		file2Name string
	}{
		{
			name:      "JSON files",
			file1Name: JSON1,
			file2Name: JSON2,
		},
		{
			name:      "YAML files",
			file1Name: YAML1,
			file2Name: YAML2,
		},
	}

	for _, tt := range tests {
		// t.Run запускает подтест для каждого случая
		t.Run(tt.name, func(t *testing.T) {
			file1Path := filepath.Join("testdata", tt.file1Name)
			file2Path := filepath.Join("testdata", tt.file2Name)

			result, err := GenDiff(file1Path, file2Path, "stylish")

			require.NoError(t, err, ERRMSG)

			assert.Equal(t, expected, result)
			assert.NotEmpty(t, result)
			assert.Contains(t, result, "host: hexlet.io")
			assert.Contains(t, result, "- follow: false")
			assert.Contains(t, result, "+ verbose: true")
		})
	}
}

func TestGenDiffIntegrationComplex(t *testing.T) {
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

	// Пути к JSON файлам
	file1Path := filepath.Join("testdata", "fixture", JSON1)

	file2Path := filepath.Join("testdata", "fixture", JSON2)

	// Генерация diff
	result, err := GenDiff(file1Path, file2Path, "stylish")

	require.NoError(t, err, ERRMSG)

	// Проверка
	assert.Equal(t, expected, result, "Сгенерированный Diff не соответствует ожидаемому")

	// Пути к YAML файлам
	file1Path = filepath.Join("testdata", "fixture", YAML1)
	file2Path = filepath.Join("testdata", "fixture", YAML2)

	// Генерация diff
	result, err = GenDiff(file1Path, file2Path, "stylish")

	require.NoError(t, err, ERRMSG)

	// Проверка
	assert.Equal(t, expected, result, "Сгенерированный Diff не соответствует ожидаемому.")
}
