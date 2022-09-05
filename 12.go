package main

import "fmt"

func makeSet12(slice []string) []string {
	// Эта функция просто убирает дубликаты из слайса, следовательно, он попадает под определение множества

	var result []string
	set := make(map[string]bool)

	for _, item := range slice {
		set[item] = true
	}

	for key, _ := range set {
		result = append(result, key)
	}

	return result
}

func main() {
	almostSet := []string{"cat", "cat", "dog", "tree", "cat"}
	fmt.Println(makeSet12(almostSet))
}
