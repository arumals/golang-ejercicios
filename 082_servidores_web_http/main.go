package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola Mundo"))
	})
	http.ListenAndServe(":8000", nil)
}
