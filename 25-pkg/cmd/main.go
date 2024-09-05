package main

import (
	"fmt"

	"github.com/garciawell/go-full-pos/pkg/math"
)

func main() {
	fmt.Println("Hello")

	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
