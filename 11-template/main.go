package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	courses := Courses{{"GO", 40}, {"JAVA", 80}}

	t := template.Must(template.New("template.html").ParseFiles("./11-template/template.html"))
	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}
}
