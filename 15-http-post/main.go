package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	content := `{"name": "Garcia"}`
	jsonVar := bytes.NewBuffer([]byte(content))
	resp, err := http.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
