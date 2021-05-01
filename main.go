package main

import (
	"bytes"
	"compress/gzip"
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

	"github.com/wcharczuk/go-chart/v2"
)

type Server struct {
	Mode     Mode
	ReadFile func(path string) ([]byte, error)
	Open     func(path string) (fs.File, error)
	Version  string
}

func (s *Server) setCache(w http.ResponseWriter) {
	switch s.Mode {
	case Dev:
		w.Header().Set("Cache-Control", "no-store")
	case Prod:
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", (time.Hour*24*7)/time.Second))
	}
}

func (s *Server) setHeader(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-type", contentType)
	w.Header().Set("Content-language", "de-DE")
	w.Header().Set("X-Robots-Tag", "noindex")
}

func (s *Server) handleFile(path, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeader(w, contentType)
		s.setCache(w)

		f, err := s.Open(path)
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
		defer f.Close()

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			io.Copy(w, f)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		io.Copy(gz, f)
	}
}

func (s *Server) handleTemplate(path, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeader(w, contentType)
		s.setCache(w)

		index, err := s.ReadFile(path)
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
			return
		}

		fm := template.FuncMap{
			"join": func(in []string) template.HTML { return template.HTML(strings.Join(in, " ")) },
		}
		templ, err := template.New("").Funcs(fm).Parse(string(index))
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
			return
		}

		data, err := s.templateData()
		if err != nil {
			http.Error(w, "nope", http.StatusNotFound)
			return
		}

		if err := templ.Execute(w, data); err != nil {
			http.Error(w, "nope", http.StatusNotFound)
		}
	}
}

func (s *Server) templateData() (interface{}, error) {
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

	usp, err := s.ReadFile("usp.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read usp.html: %w", err)
	}

	statistics, err := s.ReadFile("statistics.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read fakten.html: %w", err)
	}

	templ, err := template.New("").Parse(string(statistics))
	if err != nil {
		return nil, fmt.Errorf("error parsing statistics template: %w", err)
	}

	buf := &bytes.Buffer{}
	if err := templ.Execute(buf, struct{ Version string }{Version: s.Version}); err != nil {
		return nil, fmt.Errorf("error executing statistics template: %w", err)
	}

	data := struct {
		Month      []Month
		Version    string
		USP        template.HTML
		Statistics template.HTML
	}{
		Month:      calendar(bookings, time.March, time.April, time.May),
		Version:    s.Version,
		USP:        template.HTML(string(usp)),
		Statistics: template.HTML(buf.String()),
	}

	return data, nil
}

func (s *Server) handleChart(plot chart.StackedBarChart) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeader(w, "image/png")
		s.setCache(w)

		if err := plot.Render(chart.PNG, w); err != nil {
			http.Error(w, "nope", http.StatusInternalServerError)
			return
		}
	}
}

type Mode string

//go:embed index.html style.css favicon.ico favicon.png impressum.html car_replacement_statistics.webp usp.html
var embedded embed.FS

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

func NewServer(mode Mode, version string) *Server {
	switch mode {
	case Prod:
		return &Server{
			Mode:     mode,
			ReadFile: embedded.ReadFile,
			Open:     embedded.Open,
			Version:  version,
		}
	case Dev:
		return &Server{
			Mode:     mode,
			ReadFile: os.ReadFile,
			Open: func(path string) (fs.File, error) {
				return os.Open(path)
			},
			Version: version,
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
			log.Fatalf("error grabbing certificate: %v", err)
		}
	}

	server := NewServer(Mode(mode), version)

	handler := http.NewServeMux()
	handler.HandleFunc("/", server.handleTemplate("index.html", "text/html; charset=UTF-8"))
	handler.HandleFunc("/impressum", server.handleTemplate("impressum.html", "text/html; charset=UTF-8"))
	handler.HandleFunc("/favicon.ico", server.handleFile("favicon.ico", "image/x-icon"))
	handler.HandleFunc("/car_replacement_statistics.webp", server.handleFile("car_replacement_statistics.webp", "image/webp"))
	handler.HandleFunc(fmt.Sprintf("/chart-a-%s.png", version), server.handleChart(plotSurveyA()))
	handler.HandleFunc(fmt.Sprintf("/chart-b-%s.png", version), server.handleChart(plotSurveyB()))
	handler.HandleFunc(fmt.Sprintf("/style-%s.css", version), server.handleFile("style.css", "text/css"))
	handler.HandleFunc(fmt.Sprintf("/favicon-%s.png", version), server.handleFile("favicon.png", "image/png"))

	if Mode(mode) == Prod {
		go func() {
			http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://veedelvelo.de", http.StatusTemporaryRedirect)
			}))
		}()

		if err := http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", handler); err != nil {
			log.Fatalf("error starting web server: %v", err)
		}
	}

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("error starting web server: %v", err)
	}
}
