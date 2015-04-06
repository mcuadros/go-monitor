package monitor

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"gopkg.in/mcuadros/go-monitor.v1/aspects"
)

type Monitor struct {
	Addr    string
	Aspects map[string]aspects.Aspect

	server *http.Server
	mux    *http.ServeMux

	sync.Mutex
}

//NewMonitor returns a new Monitor, with the standard Aspects (time, runtime and memory)
func NewMonitor(addr string) *Monitor {
	m := NewPlainMonitor(addr)
	m.AddAspect(aspects.NewTimeAspect(true))
	m.AddAspect(aspects.NewRuntimeAspect(true))
	m.AddAspect(aspects.NewMemoryAspect(true))

	return m
}

//NewPlainMonitor returns a new Monitor, without aspects
func NewPlainMonitor(addr string) *Monitor {
	return &Monitor{
		Addr:    addr,
		Aspects: make(map[string]aspects.Aspect, 0),
	}
}

//AddAspect adds a new `aspects.Aspect` to the Monitor
func (m *Monitor) AddAspect(a aspects.Aspect) {
	m.Aspects[a.Name()] = a
}

//Start launch the HTTP server
func (m *Monitor) Start() error {
	m.buildServer()
	return m.server.ListenAndServe()
}

func (m *Monitor) buildServer() {
	m.server = &http.Server{Addr: m.Addr, Handler: m}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	if r.URL.Path == "/" {
		m.rootHandler(w, r)
		return
	}

	m.aspectHandler(w, r)
}

func (m *Monitor) rootHandler(w http.ResponseWriter, r *http.Request) {
	m.jsonHandle(m.getAllAspectsResults(), w, r)
}

func (m *Monitor) aspectHandler(w http.ResponseWriter, r *http.Request) {
	m.jsonHandle(m.getAspectsResults(r.URL.Path[1:]), w, r)
}

func (m *Monitor) jsonHandle(data interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (m *Monitor) getAllAspectsResults() map[string]interface{} {
	res := make(map[string]interface{}, 0)
	for k, a := range m.Aspects {
		if a.InRoot() {
			res[k] = a.GetStats()
		}
	}

	return res
}

func (m *Monitor) getAspectsResults(aspects string) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	for _, name := range strings.Split(aspects, ",") {
		if a, ok := m.Aspects[name]; ok {
			res[name] = a.GetStats()
		}
	}

	return res
}
