package main

import "fmt"

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return num
}

func main() {
	fmt.Println("5 factorial: ", factorial(5))
	fmt.Println("4 factorial: ", factorial(4))
	fmt.Println("3 factorial: ", factorial(3))
	fmt.Println("2 factorial: ", factorial(2))
	fmt.Println("1 factorial: ", factorial(1))
	fmt.Println("0 factorial: ", factorial(0))
}
