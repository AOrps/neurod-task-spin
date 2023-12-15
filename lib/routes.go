package lib

/*
   Files that Handles Routes and Function Handlers for the Webserver

*/

import (
	"fmt"
	"net/http"
	"html/template"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob(TEMPLATEPATH))
	page := Page{
		Nav: Navbar(),
	}
	tpl.ExecuteTemplate(w, "main", page)
}

func dbHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "db here plz")
}

func settingsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob(TEMPLATEPATH))
	page := Page{
		Nav: Navbar(),
	}
	tpl.ExecuteTemplate(w, "settings", page)
}

func SetupRoutes() {
	http.HandleFunc("/",rootHandle)
	http.HandleFunc("/db",dbHandle)
	http.HandleFunc("/settings",settingsHandle)	
}
