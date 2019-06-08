package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case "GET":
			w.Write([]byte("Peticion mediante GET"))
		case "POST":
			w.Write([]byte("Peticion mediante POST"))
		case "PUT":
			w.Write([]byte("Peticion mediante PUT"))
		case "DELETE":
			w.Write([]byte("Peticion mediante DELETE"))
		case "PATCH":
			w.Write([]byte("Peticion mediante PATCH"))
		default:
			w.Write([]byte("Peticion mediante metodo desconocido"))

		}
	})
	http.ListenAndServe(":8000", nil)
}
