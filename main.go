package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/todo/", todoHandlerWithID)

	fmt.Printf("Server is running...!\n")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
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

func todoHandlerWithID(w http.ResponseWriter, r *http.Request) {
	query := pathSlice("/todo/", r.URL.Path)

	id, err := strconv.Atoi(query)
	if err != nil {
		fmt.Fprintln(w, "The ID has to be a number.")
		return
	}

	todo, ok := todoTable[id]
	if !ok {
		fmt.Fprintln(w, "Please specify your registered ID")
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%d\t%s\n", id, todo)
	case "PUT":
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		todoTable[id] = body
		fmt.Fprintln(w, "success")
	case "DELETE":
		delete(todoTable, id)
		fmt.Fprintln(w, "success")
	}
}

func pathSlice(pattern string, path string) string {
	s := path[len(pattern):]
	return s
}
