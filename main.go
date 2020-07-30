package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var todoTable map[int]string
var id int

func init() {
	todoTable = make(map[int]string)
	id = 0
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/todo", todoHandler)

	fmt.Printf("Server is running...!\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is running...!")
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for i, todo := range todoTable {
			fmt.Fprintf(w, "%d\t%s\n", i, todo)
		}
	case "POST":
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		todoTable[id] = body
		id++
		fmt.Fprint(w, "Success")
	}
}
