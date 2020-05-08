package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
	Picture   string
	Bio       string
}

type Skill struct {
	Title    string
	Category string
	Level    int
	Icon     string
}

type ResumePage struct {
	Title  string
	Person Person
	Skills []Skill
}

var data = ResumePage{
	Title: "Resume",
	Person: Person{
		FirstName: "Zoldik",
		LastName:  "Zeno",
		Picture:   "zeno.png",
		Bio:       "Software engineer",
	},
	Skills: []Skill{
		{Title: "Python", Category: "lang", Level: 90, Icon: "python.png"},
		{Title: "Go", Category: "lang", Level: 60, Icon: "go.png"},
		{Title: "Java", Category: "lang", Level: 95, Icon: "java.png"},
		{Title: "Angular", Category: "lang", Level: 80, Icon: "angular.png"},
		{Title: "React", Category: "lang", Level: 55, Icon: "react.png"},
	},
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/resume-template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	fmt.Printf("Serving and listening on port %d", 8082)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8082", nil)
}
