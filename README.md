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

Утилита умеет сравнивать как плоские так и вложенные файлы. 

Алгоритм работы:
Сортировщик (BuildDiff, Генератор Дерева): Принимает на вход две кучи данных (data1 и data2) и создаёт структурированный отчёт о различиях (MapDiff). Он не знает, как этот отчёт будет выглядеть в итоге.

Форматер (StylishFormatter): Принимает этот структурированный отчёт и превращает его в красивую строку, используя нужные отступы, символы +/- и фигурные скобки.
