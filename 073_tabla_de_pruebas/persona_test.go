package persona

import "testing"

type EsAdultoTest struct {
	edad int
	want bool
}

var esAdultoTests = []EsAdultoTest{
	{edad: 17, want: false},
	{edad: 18, want: true},
	{edad: 19, want: true},
}

func TestEsAdulto(t *testing.T) {
	for _, test := range esAdultoTests {
		got := EsAdulto(test.edad)
		if got != test.want {
			t.Errorf("EsAdulto(%v) = %v, se esperaba %v", test.edad, got, test.want)
		}
	}
}
