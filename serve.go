package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	lib "github.com/AOrps/neurod-task-spin/lib"
)

const (
	templatePath = "templates/*htmx"
)

var port = 7100

// TODO: Test if PORT is available


func rootHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob(templatePath))

	page := lib.Page{
		Nav: lib.Navbar(),
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func dbHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "db here plz")
}

func main() {

	servePort := strconv.Itoa(port)

	// Assets at static/
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/",rootHandle)
	http.HandleFunc("/db",dbHandle)
	
	fmt.Printf("server: http://127.0.0.1:%s\n", servePort)
	for http.ListenAndServe(":"+servePort,nil) != nil {
		port++
		servePort = strconv.Itoa(port)
	}

}
