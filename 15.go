package main

import "fmt"

var justString string

func createHugeString(count int) string {
	text := ""
	for i := 0; i < count; i++ {
		text += "j"
	}

	return text
}

func someFunc(separate int) {
	v := createHugeString(1 << 10)

	if len(v) <= separate {
		separate = len(v)
	}

	justString = v[:separate]
	// justString = v[:100] опасно, потому что может запаниковать  slice bounds out of range
	fmt.Println(justString)
}

func main() {
	someFunc(100)
}
