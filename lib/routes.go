package lib

/*
   Files that Handles Routes and Function Handlers for the Webserver

*/

import (
	"fmt"
	"net/http"
	"html/template"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

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

func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("try: <%s> \n", string(message))
		fmt.Printf("mt: <%d>\n", mt)

		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
		}
			
	}
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
	http.HandleFunc("/ws",wsHandle)
	http.HandleFunc("/settings",settingsHandle)
}
