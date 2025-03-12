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
		Title:   "Go Subtitle Translator",
		Heading: "Translate your SRT subtitles",
	}

	wd, _ := os.Getwd()

	tmpl, err := template.ParseFiles(wd + "/templates/homepage.html")
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
