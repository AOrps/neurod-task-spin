package main

import (
	"fmt"
	"net/http"
	"strconv"
	lib "github.com/AOrps/neurod-task-spin/lib"
)

const (
	templatePath = lib.TEMPLATEPATH
)

// (Global) port: initial port to serve content on
// - should this fail, it will increment value
var port = 7100


func main() {

	// Assets at assets/
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Is the area where the routes are setup
	lib.SetupRoutes()

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
