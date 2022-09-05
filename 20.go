package main

import (
	"fmt"
	"strings"
)

func reverseSentence(input string) string {
	var result string
	stringsSlice := strings.Split(input, " ")
	stringsSliceLen := len(stringsSlice)

	for i, j := 0, stringsSliceLen-1; i < j; i, j = i+1, j-1 {
		// выглядит страшно, но i инкриментируется с нуля до конца слайса
		// а j до нуля из конца слайса

		stringsSlice[i], stringsSlice[j] = stringsSlice[j], stringsSlice[i]
	}

	for _, s := range stringsSlice {
		result += s + " "
	}

	return strings.TrimSuffix(result, " ")
}

func main() {
	before := "snow dog sun"
	after := reverseSentence(before)
	fmt.Printf("%v — %v\n", before, after)
	fmt.Println(len(before) == len(after))
}
