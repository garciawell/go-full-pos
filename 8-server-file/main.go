package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./8-server-file/public"))
	muxFile := http.NewServeMux()
	muxFile.Handle("/", fileServer)
	fmt.Print("Running 8080...")
	log.Fatal(http.ListenAndServe(":8080", muxFile))
}
