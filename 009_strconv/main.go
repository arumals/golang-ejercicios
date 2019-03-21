package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var mayorDeEdad bool = true
	strVal := strconv.FormatBool(mayorDeEdad)
	fmt.Println(strVal, reflect.TypeOf(strVal))
}
