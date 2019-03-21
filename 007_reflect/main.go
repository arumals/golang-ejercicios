package main

import (
	"fmt"
	"reflect"
)

func main() {
	var alumno string = "Jose Luis"
	var edad int = 22
	var peso float64 = 85.5
	fmt.Println(reflect.TypeOf(alumno))
	fmt.Println(reflect.TypeOf(edad))
	fmt.Println(reflect.TypeOf(peso))
}
