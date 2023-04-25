package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}
	
	u := User{
		Name: "Oskar",
	}

	err = t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}

}