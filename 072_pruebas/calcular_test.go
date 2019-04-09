package calcular

import "testing"

func TestSumar(t *testing.T) {
	if Sumar(2, 2) != 4 {
		t.Fatal("La suma de 2 + 2 debe ser igual a 4")
	}
}

func TestRestar(t *testing.T) {
	if Restar(5, 3) != 2 {
		t.Fatal("La resta de 5 - 3 debe ser igual a 2")
	}
}

func TestMultiplicar(t *testing.T) {
	if Multiplicar(4, 4) != 16 {
		t.Fatal("La multiplicación de 4 x 4 debe ser igual a 16")
	}
}

func TestDividir(t *testing.T) {
	if Dividir(10, 2) != 5 {
		t.Fatal("La división de 10 / 2 debe ser igual a 5")
	}
}
