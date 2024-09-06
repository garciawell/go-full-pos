package main

import (
	"fmt"

	"github.com/garciawell/go-full-pos/apis/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
