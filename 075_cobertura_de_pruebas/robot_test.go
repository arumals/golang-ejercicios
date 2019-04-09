package robot

import "testing"

type ControladorTest struct {
	codigo int
	want   string
}

var controladorTests = []ControladorTest{
	{codigo: 100, want: "encender"},
	{codigo: 200, want: "apagar"},
	{codigo: 300, want: "avanzar"},
	{codigo: 400, want: "saludar"},
	{codigo: 500, want: ""},
}

func TestControladorTests(t *testing.T) {
	for _, test := range controladorTests {
		got, _ := Controlador(test.codigo)
		if got != test.want {
			t.Errorf("Controlador(%v) = %v, se esperaba %v", test.codigo, got, test.want)
		}
	}
}
