package main

import (
	"fmt"
	"reflect"
)

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return 3.1416 * c.Radio * c.Radio
}

func (c *Circulo) Perimetro() float64 {
	return 2 * 3.1416 * c.Radio
}

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c *Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}

func ImprimirDetalles(f Figura) {
	fmt.Println("El área es: ", reflect.TypeOf(f).Elem().Name(), f.Area())
	fmt.Println("El perímetro es: ", reflect.TypeOf(f).Elem().Name(), f.Perimetro())
}

func main() {
	figura1 := Circulo{Radio: 5.0}
	figura2 := Cuadrado{Lado: 8.0}
	ImprimirDetalles(&figura1)
	ImprimirDetalles(&figura2)
}
