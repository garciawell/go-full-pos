package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./file/arquivo.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Hello World"))

	if err != nil {
		panic(err)
	}
	f.Close()

	fmt.Printf("O tamanho do arquivo é %d bytes \n", size)

	/// READ
	file, err := os.ReadFile("./file/arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	file2, err := os.Open("./file/arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	os.Remove("./file/arquivo.txt")
}
