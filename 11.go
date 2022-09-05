package main

import "fmt"

func makeSet(slice []int) []int {
	var result []int
	set := make(map[int]bool)

	for _, item := range slice {
		set[item] = true
	}

	for key, _ := range set {
		result = append(result, key)
	}

	return result
}

func main() {
	nums1 := makeSet([]int{23, 42, 11, 211, 11, 534, 42})
	nums2 := makeSet([]int{34, 232, 42})

	var biggest []int  // эти два слайса нужны для опеределения, по какому из них сначала надо итерироваться
	var smallest []int // эти два слайса нужны для опеределения, по какому из них сначала надо итерироваться
	var commonNums []int

	if len(nums1) > len(nums2) {
		// определяем, по какому из них будем итерироваться в первую очередь
		biggest = nums1
		smallest = nums2
	} else {
		biggest = nums2
		smallest = nums1
	}

	for _, num1 := range smallest {
		// итерируемся сначала по меньшему, чтобы сравнить все числа
		for _, num2 := range biggest {
			if num1 == num2 {
				commonNums = append(commonNums, num1)
			}
		}
	}

	fmt.Println(commonNums)
}
