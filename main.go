package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := loadTemplates()
		fileName := r.URL.Path[1:]
		t := templates.Lookup(fileName)
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}else {
			w.WriteHeader(http.StatusNotFound)
		}
	});
	http.Handle("/assets/", http.FileServer(http.Dir("templates")))
	http.Handle("/src/assets/img/", http.FileServer(http.Dir("templates")))
	http.Handle("/dependencies/", http.FileServer(http.Dir("templates")))

	http.ListenAndServe("localhost:8080", nil)
}

func loadTemplates() *template.Template {
	result := template.New("templates")
	result, err := result.ParseGlob("templates/*.html")

	template.Must(result, err)
	return result
}
