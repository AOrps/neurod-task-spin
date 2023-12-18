package lib

/*
   Files that Handles Routes and Function Handlers for the Webserver

*/

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var todoList []string

func getCmd(input string) string {
    inputArr := strings.Split(input, " ")
    return inputArr[0]
}

func getMessage(input string) string {
    inputArr := strings.Split(input, " ")
    var result string
    for i := 1; i < len(inputArr); i++ {
        result += inputArr[i]
    }
    return result
}

func updateTodoList(input string) {
    tmpList := todoList
    todoList = []string{}
    for _, val := range tmpList {
        if val == input {
            continue
        }
        todoList = append(todoList, val)
    }
}


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

		input := string(message)
		cmd := getCmd(input)
		msg := getMessage(input)

		if cmd == "add" {
                todoList = append(todoList, msg)
            } else if cmd == "done" {
                updateTodoList(msg)
            }
            output := "Current Todos: \n"
            for _, todo := range todoList {
                output += "\n - " + todo + "\n"
            }
            output += "\n----------------------------------------"
            message = []byte(output)
            err = conn.WriteMessage(mt, message)
            if err != nil {
                fmt.Println("write failed:", err)
                break
            }
	
	}
}

func wsPage(w http.ResponseWriter, r *http.Request) {
	// tpl := template.Must(template.ParseGlob("templates/*html"))
	// tpl.ExecuteTemplate(w, "settings", nil)
	http.ServeFile(w, r, "templates/ws-try.html")


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
	http.HandleFunc("/test",wsPage)
	http.HandleFunc("/ws",wsHandle)
	http.HandleFunc("/settings",settingsHandle)
}
