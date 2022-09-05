package main

import "fmt"

func convertToRunes(str string) []rune {
	var runes []rune

	for _, char := range str {
		runes = append(runes, char)
	}

	return runes
}

func isUnique(str string) bool {

	runes := convertToRunes(str)

	keys := make(map[rune]bool)
	list := []rune{}

	for _, entry := range runes {
		_, value := keys[entry]

		if !value {
			// Дубликатов не найдено, добавляем в список уже известных рун

			keys[entry] = true
			list = append(list, entry)
		} else {
			// Найден дубликат
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isUnique("aA"))
}
