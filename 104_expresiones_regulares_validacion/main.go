package main

import (
	"fmt"
	"regexp"
)

func main() {
	// expresion regular
	re := regexp.MustCompile("^[0-9]{4}(-[0-9]{2}){2} [0-9]{2}(:[0-9]{2}){2}$")
	// validar diferentes valores
	fmt.Println(re.MatchString("19-01-01 00:00:00"))
	fmt.Println(re.MatchString("2019-01-01 00:00:AA"))
	fmt.Println(re.MatchString("2019-01-01 00:00:00"))
}
