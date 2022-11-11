// my first project with html and go very odd combo.
//
// content sources so far:
// https://archive.org/details/southparksquidgames
// https://archive.org/details/family-guy_202204
//
// the guts of this project:
// https://freshman.tech/web-development-with-go/
// 

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/joho/godotenv"
)
var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>ZER0x Server | In development.</h1>"))
	//w.Write([]byte("<h2></h2>"))
	tpl.Execute(w, nil)
}


func main() {
	fmt.Println("Server started (LOCAL)")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}