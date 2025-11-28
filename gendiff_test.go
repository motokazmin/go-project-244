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

			result := GenDiff(data1, data2)
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

			result := GenDiff(data1, data2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenDiffIntegration(t *testing.T) {
	// Используем файлы из testdata
	file1Path := filepath.Join("testdata", "file1.json")
	file2Path := filepath.Join("testdata", "file2.json")

	// Парсим файлы
	data1, err := ParseFile(file1Path)
	require.NoError(t, err)

	data2, err := ParseFile(file2Path)
	require.NoError(t, err)

	// Генерируем diff
	result := GenDiff(data1, data2)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	assert.Equal(t, expected, result)

	// Проверяем что результат не пустой
	assert.NotEmpty(t, result)

	// Проверяем что содержит ключевые элементы
	assert.Contains(t, result, "host: hexlet.io")
	assert.Contains(t, result, "- follow: false")
	assert.Contains(t, result, "+ verbose: true")
}

func TestGenDiffIntegrationYaml(t *testing.T) {
	// Используем файлы из testdata
	file1Path := filepath.Join("testdata", "file1.yaml")
	file2Path := filepath.Join("testdata", "file2.yaml")

	// Парсим файлы
	data1, err := ParseFile(file1Path)
	require.NoError(t, err)

	data2, err := ParseFile(file2Path)
	require.NoError(t, err)

	// Генерируем diff
	result := GenDiff(data1, data2)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	assert.Equal(t, expected, result)

	// Проверяем что результат не пустой
	assert.NotEmpty(t, result)

	// Проверяем что содержит ключевые элементы
	assert.Contains(t, result, "host: hexlet.io")
	assert.Contains(t, result, "- follow: false")
	assert.Contains(t, result, "+ verbose: true")
}
