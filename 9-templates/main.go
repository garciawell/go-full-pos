package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"GO", 40}
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}} - Carga Horária {{.Workload}}")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
