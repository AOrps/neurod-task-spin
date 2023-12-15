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

// (Global) port: initial port to serve content on
// - should this fail, it will increment value
var port = 7100


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

	// Assets at static/
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/",rootHandle)
	http.HandleFunc("/db",dbHandle)

	// checks if port is available and valid
	for (lib.CheckPort(port) == false) && (port < 65536) {
		port++
	}

	// converts port number to string (so ListenAndServe)
	// can use it post conversion
	servePort := strconv.Itoa(port)
	
	fmt.Printf("server: http://127.0.0.1:%s\n", servePort)
	http.ListenAndServe(":"+servePort,nil)
}
