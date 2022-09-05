package main

import "fmt"

func deleteItemByIndex(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	result := deleteItemByIndex(slice1, 2)
	fmt.Println(result)
}
