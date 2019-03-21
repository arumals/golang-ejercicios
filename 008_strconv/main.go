package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var mayorDeEdad string = "true"
	boolVal, _ := strconv.ParseBool(mayorDeEdad)
	fmt.Println(boolVal, reflect.TypeOf(boolVal))
}
