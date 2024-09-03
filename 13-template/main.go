package main

import (
	"net/http"
	"strings"
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

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Server(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"13-template/header.html",
		"13-template/content.html",
		"13-template/footer.html",
	}

	courses := Courses{{"Go", 40}, {"java", 80}}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(w, courses)
	if err != nil {
		panic(err)
	}
}
