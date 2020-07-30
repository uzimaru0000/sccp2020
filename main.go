package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/uzimaru0000/kapro/repo"
	"github.com/uzimaru0000/kapro/todo"
)

var todoRepo repo.TodoRepo

func init() {
	todoRepo = repo.NewInMemTodoRepo()
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/todo", todoHandler)
	http.HandleFunc("/todo/", todoHandlerWithID)

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
		todoList, err := todoRepo.GetAll()
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}

		for _, todo := range todoList {
			fmt.Fprintf(w, "%s\n", todo.ToString())
		}
	case "POST":
		body := getBody(r)
		err := todoRepo.Store(todo.New(body))
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		fmt.Fprint(w, "Success")
	}
}

func todoHandlerWithID(w http.ResponseWriter, r *http.Request) {
	id := pathSlice("/todo/", r.URL.Path)

	todo, err := todoRepo.Get(&todo.Todo{Id: id})
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%s\n", todo.ToString())
	case "PUT":
		body := getBody(r)
		todo.Update(body)
		todoRepo.Update(todo)
		fmt.Fprintln(w, "success")
	case "DELETE":
		todoRepo.Delete(todo)
		fmt.Fprintln(w, "success")
	}
}

func getBody(r *http.Request) string {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	return bufbody.String()
}

func pathSlice(pattern string, path string) string {
	s := path[len(pattern):]
	return s
}
