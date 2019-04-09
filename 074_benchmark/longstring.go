package longstring

import (
	"bytes"
	"strings"
)

func MedianteConcatenacion(longitud int) string {
	var s string
	for i := 0; i < longitud; i++ {
		s += "texto"
	}
	return s
}

func MedianteArreglo(longitud int) string {
	s := []string{}
	for i := 0; i < longitud; i++ {
		s = append(s, "texto")
	}
	return strings.Join(s, "")
}

func MedianteBuffer(longitud int) string {
	var buffer bytes.Buffer
	for i := 0; i < longitud; i++ {
		buffer.WriteString("texto")
	}
	return buffer.String()
}
