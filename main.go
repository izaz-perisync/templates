package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Dogs struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

func main() {
	// r := mux.NewRouter()
	// api := r.PathPrefix("/api").Subrouter()

	// template := api.PathPrefix("/template").Subrouter()

	dogs := []Dogs{
		{
			Name:   "<script>alert(\"Gotcha!\");</script>Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pit Bull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
		{
			Name:   "Bruce Wayne",
			Sex:    "Male",
			Intact: false,
			Age:    "3 years, 8 months",
			Breed:  "Chihuahua",
		},
	}

	// var tmptFile = "petsHtml.tmpl"
	// tmpl, err := template.New(tmptFile).ParseFiles(tmptFile)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// . . .
	// funcMap := template.FuncMap{
	// 	"dec":     func(i int) int { return i - 1 },
	// 	"replace": strings.ReplaceAll,
	// 	"join":strings.Join,
	// }
	// // var tmplFile = “lastPet.tmpl”
	// tmpl, err := template.New(tmptFile).Funcs(funcMap).ParseFiles(tmptFile)
	// if err != nil {
	// 	panic(err)
	// }

	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
	}
	var tmplFile = "petsHtml.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	var f *os.File
	f, err = os.Create("pets.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, dogs)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	// end main
	// . . .
	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		fmt.Println("err", err)
	}

}
