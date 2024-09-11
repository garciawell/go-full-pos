package main

import "fmt"

func recebe(name string, hello chan<- string) {
	hello <- "Hello " + name
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Gopher", hello)
	ler(hello)
}
