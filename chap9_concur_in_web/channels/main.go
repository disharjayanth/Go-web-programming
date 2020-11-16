package main

import (
	"fmt"
)

func callerA(a chan<- string) {
	a <- "Hello World!"
	close(a)
}

func callerB(b chan<- string) {
	b <- "Hola Mundo!"
	close(b)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from A\n", msg)
			}
		}
	}
}
