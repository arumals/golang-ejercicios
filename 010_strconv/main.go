package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var edad string = "50"
	numVal, _ := strconv.ParseInt(edad, 10, 64)
	fmt.Println(numVal, reflect.TypeOf(numVal))
}
