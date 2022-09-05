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

	// type assertion. Можно сделать ещё проверку по типу и запаниковать, если он не подходит
	variableInt, isInt := variable.(int)
	fmt.Println(variableInt, isInt)
}
