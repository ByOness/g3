package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Fail parse template")
		return
	}
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./content/")
	staticFileHandler := http.StripPrefix("/content/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/content/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/g3", MainHandler)
	http.Handle("/", r)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8082", nil)
}
