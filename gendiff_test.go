package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name      string
		file1Path string
		file2Path string
		expected  string
	}{
		{
			name:      "main example",
			file1Path: filepath.Join("testdata", "file1.json"),
			file2Path: filepath.Join("testdata", "file2.json"),
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
			data1, err := ParseFile(tt.file1Path)
			require.NoError(t, err)

			data2, err := ParseFile(tt.file2Path)
			require.NoError(t, err)

			result := GenDiff(data1, data2, "stylish")
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
			file1Path: filepath.Join("testdata", "file1.json"),
			file2Path: filepath.Join("testdata", "file1.json"),
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
			data1, err := ParseFile(tt.file1Path)
			require.NoError(t, err)

			data2, err := ParseFile(tt.file2Path)
			require.NoError(t, err)

			result := GenDiff(data1, data2, "stylish")
			assert.Equal(t, tt.expected, result)
		})
	}
}

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
			file1Name: "file1.json",
			file2Name: "file2.json",
		},
		{
			name:      "YAML files",
			file1Name: "file1.yaml",
			file2Name: "file2.yaml",
		},
	}

	for _, tt := range tests {
		// t.Run запускает подтест для каждого случая
		t.Run(tt.name, func(t *testing.T) {
			file1Path := filepath.Join("testdata", tt.file1Name)
			file2Path := filepath.Join("testdata", tt.file2Name)

			data1, err := ParseFile(file1Path)
			require.NoError(t, err)

			data2, err := ParseFile(file2Path)
			require.NoError(t, err)

			result := GenDiff(data1, data2, "stylish")

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

	// 1. Пути к JSON файлам
	file1Path := filepath.Join("testdata", "fixture", "file1.json")
	file2Path := filepath.Join("testdata", "fixture", "file2.json")

	// 2. Парсинг файлов
	data1, err := ParseFile(file1Path)
	require.NoError(t, err, "Ошибка при парсинге %s. Убедитесь, что файлы существуют.", file1Path)

	data2, err := ParseFile(file2Path)
	require.NoError(t, err, "Ошибка при парсинге %s. Убедитесь, что файлы существуют.", file2Path)

	// 3. Генерация diff
	result := GenDiff(data1, data2, "stylish")

	// 4. Проверка
	assert.Equal(t, expected, result, "Сгенерированный Diff не соответствует ожидаемому.")

	// 1. Пути к YAML файлам
	file1Path = filepath.Join("testdata", "fixture", "file1.yaml")
	file2Path = filepath.Join("testdata", "fixture", "file2.yaml")

	// 2. Парсинг файлов
	data1, err = ParseFile(file1Path)
	require.NoError(t, err, "Ошибка при парсинге %s. Убедитесь, что файлы существуют.", file1Path)

	data2, err = ParseFile(file2Path)
	require.NoError(t, err, "Ошибка при парсинге %s. Убедитесь, что файлы существуют.", file2Path)

	// 3. Генерация diff
	result = GenDiff(data1, data2, "stylish")

	// 4. Проверка
	assert.Equal(t, expected, result, "Сгенерированный Diff не соответствует ожидаемому.")
}
