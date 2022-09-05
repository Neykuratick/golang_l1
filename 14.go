package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}
	variable = "true"
	varType := reflect.ValueOf(variable).Kind()

	fmt.Println(varType)

	// type assertion
	variableInt, isInt := variable.(int)
	fmt.Println(variableInt, isInt)
}
