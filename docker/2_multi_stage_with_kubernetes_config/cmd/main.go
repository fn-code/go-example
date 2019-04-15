package main

import (
	"log"
	"net/http"

	"html/template"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "version information about this binary",
		ConstLabels: map[string]string{
			"version": "0.0.1",
		},
	})

	httpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Count of all http request",
	}, []string{"code", "method"})
)

func home(w http.ResponseWriter, r *http.Request) {
	rs := New(w)
	rs.Name = "template"
	rs.Template = []string{"template", "home"}
	rs.Vars["data"] = "selamat malam"

	err := rs.Render()
	if err != nil {
		http.Error(w, "Failed Load template", http.StatusInternalServerError)
	}
}

func main() {

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestTotal)
	r.MustRegister(version)

	mx := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mx.Handle("/static/", http.StripPrefix("/static/", fs))
	ins := http.HandlerFunc(home)
	// mx.HandleFunc("/", home)

	mx.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestTotal, ins))
	mx.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	log.Println("Server running on port :8080")
	if err := http.ListenAndServe(":8080", mx); err != nil {
		log.Fatal(err)
	}

}

// View is for store view config
type View struct {
	request   http.ResponseWriter
	Name      string
	Template  []string
	extension string
	Vars      map[string]interface{}
	path      string
}

// New return view
func New(w http.ResponseWriter) *View {
	return &View{
		request:   w,
		extension: ".html",
		path:      "./template/",
		Vars:      make(map[string]interface{}),
	}
}

// Render is for rendering template
func (v *View) Render() error {
	var tplList []string

	for _, t := range v.Template {
		path, err := filepath.Abs(v.path + t + v.extension)
		if err != nil {
			return err
		}
		tplList = append(tplList, path)
	}
	tpl, err := template.New(v.Name).Delims("{%", "%}").ParseFiles(tplList...)
	if err != nil {
		return err
	}

	if err := tpl.ExecuteTemplate(v.request, v.Name+v.extension, v.Vars); err != nil {
		return err
	}
	return nil
}
