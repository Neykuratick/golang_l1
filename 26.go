package main

import "fmt"

func isUnique(str string) bool {
	var list []rune
	runes := []rune(str)
	keys := make(map[rune]bool)

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
