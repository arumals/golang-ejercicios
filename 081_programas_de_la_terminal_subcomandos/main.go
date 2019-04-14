package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	flag.Usage = func() {
		documentacion := `Las opciones disponibles son
mayus Convierte el texto a mayúsculas
minus Convierte el texto a minúsculas`
		fmt.Fprintf(os.Stderr, "%s\n", documentacion)
	}

	subCmdMayus := flag.NewFlagSet("mayus", flag.ExitOnError)
	subCmdMinus := flag.NewFlagSet("minus", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "mayus":
		s := subCmdMayus.String("s", "", "Introduzca el texto a convertir en mayúsculas")
		subCmdMayus.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "minus":
		s := subCmdMinus.String("s", "", "Introduzca el texto a convertir en minúsculas")
		subCmdMinus.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()
	}
}
