package front

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Heading string
	}{
		Title:   "Front Page",
		Heading: "Welcome to the front page",
	}

	wd, _ := os.Getwd()

	tmpl, err := template.ParseFiles(wd + "/templates/base.html")
	if err != nil {
		log.Fatal("Template error:", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal("Template error:", err)
		return
	}
}
