package main

import "fmt"

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return 3.1416 * c.Radio * c.Radio
}

func (c *Circulo) Perimetro() float64 {
	return 2 * 3.1416 * c.Radio
}

func main() {
	circulo1 := Circulo{Radio: 13}
	fmt.Println("El area del circulo es: ", circulo1.Area())
}
