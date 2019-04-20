package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hola Mundo"))
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
