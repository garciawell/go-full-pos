package main

import (
	"fmt"

	"github.com/garciawell/go-full-pos/25-pkg/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
