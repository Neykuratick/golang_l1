package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.44, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, num := range temperatures {
		firstDigit, _ := math.Modf(num / 10)       // получаем целую часть
		group := int(firstDigit * 10)              // преобразуем его в десятки
		groups[group] = append(groups[group], num) // добавляем число в его группу
	}

	fmt.Println(groups)
}
