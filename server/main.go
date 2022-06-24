package main

import (
	"fmt"
	"log"
	"net/http"
	"server/controller"
)

func main() {

	http.Handle("/", controller.Fileserver)
	http.HandleFunc("/form", Formhandler)
	http.HandleFunc("/hello", Hellohandler)

	fmt.Printf("Listening on port 8080. . .")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
