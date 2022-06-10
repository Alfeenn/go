package handler

import (
	"fmt"
	"net/http"
)

func Hellohandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported !", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func Formhandler(w http.ResponseWriter, r *http.Request) {

	if r.ParseForm() != nil {

		http.Error(w, "Error Parse(): %v", http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name :%s", name)
	fmt.Fprintf(w, "address :%s", address)

}
