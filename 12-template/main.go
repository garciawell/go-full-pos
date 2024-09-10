package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", Server)
	http.ListenAndServe(":8080", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	courses := Courses{{"GO", 40}, {"JAVA", 80}}
	t := template.Must(template.New("template.html").ParseFiles("./11-template/template.html"))
	err := t.Execute(w, courses)
	if err != nil {
		panic(err)
	}
}
