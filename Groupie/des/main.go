package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	parse, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("error parsing", err)
		return
	}

	parse.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error serving", err)
		return
	}
}
