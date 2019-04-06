package dimensiones

import (
	"testing"
)

func TestMillasEnKm(t *testing.T) {
	if kms := MillasEnKm(3.0); kms != 4.82802 {
		t.Errorf("se esperaba %v, se recibió %v", 4.82802, kms)
	}
}

func TestPiesEnMetros(t *testing.T) {
	if mts := PiesEnMetros(3); mts != 0.9144000000000001 {
		t.Errorf("se esperaba %v, se recibió %v", 0.9144000000000001, mts)
	}
}

func TestPulgadasEnCm(t *testing.T) {
	if cms := PulgadasEnCentimentros(3); cms != 7.62 {
		t.Errorf("se esperaba %v, se recibió %v", 7.62, cms)
	}
}
