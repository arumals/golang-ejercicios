package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		w.Write([]byte("Se realizó una petición GET"))
	case "POST":
		w.Write([]byte("Se realizó una petición POST"))
	case "PUT":
		w.Write([]byte("Se realizó una petición PUT"))
	case "DELETE":
		w.Write([]byte("Se realizó una petición DELETE"))
	default:
		w.Write([]byte("Se realizó una petición " + r.Method))
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
