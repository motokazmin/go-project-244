### Hexlet tests and linter status:
[![Actions Status](https://github.com/motokazmin/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/motokazmin/go-project-244/actions)

Gendiff
Утилита для сравнения конфигурационных файлов.
Описание
Gendiff находит различия между двумя файлами (JSON или YAML) и выводит их в удобочитаемом формате.
Установка
bashgo get github.com/urfave/cli/v3
go get gopkg.in/yaml.v3
go build -o bin/gendiff
Использование
bashbin/gendiff <filepath1> <filepath2>
Опции

-f, --format - формат вывода (по умолчанию: stylish)

