package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func handleFile(path, contentType string, header http.Header) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		w.Header().Set("Cache-Control", "no-store")
		for k, v := range header {
			for i := range v {
				w.Header().Add(k, v[i])
			}
		}
		f, err := os.Open(path)
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
		defer f.Close()
		io.Copy(w, f)
	}
}

func handleTemplate(path, contentType string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		w.Header().Set("Cache-Control", "no-store")
		index, err := os.ReadFile("index.html")
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
		fm := template.FuncMap{
			"join": func(in []string) template.HTML { return template.HTML(strings.Join(in, " ")) },
		}
		templ, err := template.New("").Funcs(fm).Parse(string(index))
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
		if err := templ.Execute(w, data); err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
	}
}

func main() {
	handler := http.NewServeMux()

	contentHeader := make(http.Header)
	contentHeader.Set("Content-language", "de-DE")
	bookings := []Booking{
		{
			From: time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			To:   time.Date(2021, time.April, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			From: time.Date(2021, time.March, 15, 0, 0, 0, 0, time.UTC),
			To:   time.Date(2021, time.March, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			From: time.Date(2021, time.May, 8, 0, 0, 0, 0, time.UTC),
			To:   time.Date(2021, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
	}

	data := calendar(bookings, time.March, time.April, time.May)

	handler.HandleFunc("/", handleTemplate("index.html", "text/html; charset=UTF-8", struct{ Month []Month }{Month: data}))
	handler.HandleFunc("/favicon.ico", handleFile("favicon.ico", "image/x-icon", make(http.Header)))
	handler.HandleFunc("/style.css", handleFile("style.css", "text/css", make(http.Header)))
	handler.HandleFunc("/favicon.png", handleFile("favicon.png", "image/png", make(http.Header)))

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("error starting web server: %w", err)
	}
}
