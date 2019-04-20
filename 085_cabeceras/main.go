package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"mensaje":"Hola Mundo"}`))
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
