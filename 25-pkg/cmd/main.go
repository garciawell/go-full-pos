package main

import (
	"fmt"

	"github.com/garciawell/go-full-pos/pkg/math"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello")

	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
