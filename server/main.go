package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", Formhandler)
	http.HandleFunc("/hello", Hellohandler)

	fmt.Printf("Listening on port 8080. . .")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}