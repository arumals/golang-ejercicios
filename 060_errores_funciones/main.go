package main

import (
	"fmt"
)

func venderCerveza(edad int, cantidad int) (int, error) {
	if edad < 18 {
		err := fmt.Errorf("No se puede vender cerveza a menores")
		return 0, err
	}
	return cantidad, nil
}

func main() {
	// variables de inicio
	var cantidad int
	var err error
	// persona de 15 años intenta comprar 6 cervezas
	cantidad, err = venderCerveza(15, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se vendieron", cantidad, "cervezas")
	}
	// persona de 21 años intenta comprar 24 cervezas
	cantidad, err = venderCerveza(21, 24)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se vendieron", cantidad, "cervezas")
	}
}
