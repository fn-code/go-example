package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data := "selamat malam"
	if err := render(w, "template", data, "./template/template.gohtml"); err != nil {
		http.Error(w, "Failed Load template", http.StatusInternalServerError)
	}
}

func render(w http.ResponseWriter, nm, data string, tplList ...string) error {
	tpl, err := template.New(nm).Delims("{%", "%}").ParseFiles(tplList...)
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(w, nm+".gohtml", data); err != nil {
		return err
	}
	return nil
}

func main() {
	mx := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mx.Handle("/static/", http.StripPrefix("/static/", fs))
	mx.HandleFunc("/", home)
	log.Println("Server running on port :8080")
	if err := http.ListenAndServe(":8080", mx); err != nil {
		log.Fatal(err)
	}
}
