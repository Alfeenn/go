package main

import (
	"fmt"
)

func main() {

	var message = make(chan string)
	go sendmsg(message)
	printmsg(message)

}

func sendmsg(msg chan<- string) {
	for i := 0; i < 10; i++ {
		msg <- fmt.Sprintf("Data %d,", i)
	}
	close(msg)
}

func printmsg(m <-chan string) {
	for message := range m {
		fmt.Println(message)
	}
}
