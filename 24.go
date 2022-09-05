package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func calculateDistance(point1 Point, point2 Point) float64 {
	result := math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2) // Нашёл формулу в учебники алгебры
	return math.Sqrt(result)
}

func main() {
	a := Point{x: -4, y: 30}
	b := Point{x: 57, y: -42}

	fmt.Println(calculateDistance(a, b))
}
