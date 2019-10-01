package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/page.html")
	tmpl.Execute(w,nil)
}

