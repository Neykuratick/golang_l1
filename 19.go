package main

import "fmt"

func reverseString(input string) string {

	chars := []rune(input)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)

}

func main() {
	fmt.Println(reverseString("главрыба"))
	fmt.Println(reverseString("шалаш"))
	fmt.Println(reverseString("кабак"))
	fmt.Println(reverseString("заказ"))
	fmt.Println(reverseString("наворован"))
}
