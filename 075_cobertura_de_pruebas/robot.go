package robot

import "fmt"

func Controlador(i int) (comando string, err error) {
	err = nil
	switch i {
	case 100:
		comando = "encender"
		break
	case 200:
		comando = "apagar"
		break
	case 300:
		comando = "avanzar"
		break
	case 400:
		comando = "saludar"
		break
	default:
		comando = ""
		err = fmt.Errorf("Opción invalida")
	}
	return
}
