package mainсщ

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	First   string
	Surname string
	Phone   string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5434", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "Registration",
		First:   "Name",
		Surname: "Surname",
		Phone:   "Phone",
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}
