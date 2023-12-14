package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	port = 7100
)

// TODO: Test if PORT is available

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yo")
}

func dbHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "db here plz")
}

func main() {

	port := strconv.Itoa(port)

	http.HandleFunc("/",rootHandle)
	http.HandleFunc("/db",dbHandle)
	
	fmt.Printf("server: http://127.0.0.1:%s\n", port)
	http.ListenAndServe(":"+port,nil)

}
