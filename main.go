package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Server struct {
	Mode     Mode
	ReadFile func(path string) ([]byte, error)
	Open     func(path string) (fs.File, error)
}

func (s *Server) setCache(w http.ResponseWriter) {
	switch s.Mode {
	case Dev:
		w.Header().Set("Cache-Control", "no-store")
	case Prod:
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", (time.Hour*24*7)/time.Second))
	}
}

func (s *Server) handleFile(path, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		s.setCache(w)

		f, err := s.Open(path)
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
		defer f.Close()
		io.Copy(w, f)
	}
}

func (s *Server) handleTemplate(path, contentType string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		w.Header().Set("Content-language", "de-DE")
		s.setCache(w)

		index, err := s.ReadFile("index.html")
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

type Mode string

//go:embed index.html style.css favicon.ico favicon.png
var embedded embed.FS

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

func NewServer(mode Mode) *Server {
	switch mode {
	case Prod:
		return &Server{
			Mode:     mode,
			ReadFile: embedded.ReadFile,
			Open:     embedded.Open,
		}
	case Dev:
		return &Server{
			Mode:     mode,
			ReadFile: os.ReadFile,
			Open: func(path string) (fs.File, error) {
				return os.Open(path)
			},
		}
	default:
		panic("operating mode not supported")
	}
}

func supportedModeP(in Mode) bool {
	for _, m := range []Mode{Dev, Prod} {
		if m == in {
			return true
		}
	}
	return false
}

func main() {
	var mode, port string
	var refereshCert bool
	var version string

	flag.StringVar(&mode, "mode", "prod", "operating mode (prod, dev)")
	flag.StringVar(&port, "port", "8080", "port number")
	flag.BoolVar(&refereshCert, "refresh-cert", false, "refresh cert")
	flag.StringVar(&version, "version", "1", "asset version")
	flag.Parse()

	if !supportedModeP(Mode(mode)) {
		log.Fatalf("mode not supported: %s", mode)
	}

	if refereshCert {
		if err := generateCert(); err != nil {
			log.Fatalf("error grabbing certificate: %w", err)
		}
	}

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

	server := NewServer(Mode(mode))

	data := struct {
		Month   []Month
		Version string
	}{
		Month:   calendar(bookings, time.March, time.April, time.May),
		Version: version,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", server.handleTemplate("index.html", "text/html; charset=UTF-8", data))
	handler.HandleFunc("/favicon.ico", server.handleFile("favicon.ico", "image/x-icon"))
	handler.HandleFunc(fmt.Sprintf("/style-%s.css", version), server.handleFile("style.css", "text/css"))
	handler.HandleFunc(fmt.Sprintf("/favicon-%s.png", version), server.handleFile("favicon.png", "image/png"))

	if Mode(mode) == Prod {
		go func() {
			http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://veedelvelo.de", http.StatusTemporaryRedirect)
			}))
		}()

		if err := http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", handler); err != nil {
			log.Fatalf("error starting web server: %w", err)
		}
	}

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("error starting web server: %w", err)
	}
}
