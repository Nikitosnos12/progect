package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handleFunc()
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5623", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")

	if err != nil {
		fmt.Fprintf(w, err.Error)
	}

	t.ExecuteTemplate(w, "idex", nil)

}
