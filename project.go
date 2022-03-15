package progect

import (
	"net/http"
)
func main () {
	handleFunc()
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/imdex.html")