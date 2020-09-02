package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	maps := map[string]string{
		"india":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}
	err := tpl.Execute(os.Stdout, maps)
	if err != nil {
		log.Fatalln(err)
	}
}
