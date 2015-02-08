package monitor

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/mcuadros/go-monitor/aspects"
)

type Monitor struct {
	Addr    string
	Aspects map[string]aspects.Aspect

	server *http.Server
	mux    *http.ServeMux

	sync.Mutex
}

func NewMonitor(addr string) *Monitor {
	m := &Monitor{
		Addr:    addr,
		Aspects: make(map[string]aspects.Aspect, 0),
	}

	m.AddAspect(&aspects.RuntimeAspect{true})
	m.AddAspect(&aspects.MemoryAspect{true})

	return m
}

func (m *Monitor) AddAspect(a aspects.Aspect) {
	m.Aspects[a.Name()] = a
}

func (m *Monitor) Start() error {
	m.buildServer()
	return m.server.ListenAndServe()
}

func (m *Monitor) buildServer() {
	m.server = &http.Server{Addr: m.Addr, Handler: m}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		m.rootHandler(w, r)
		return
	}

	m.aspectHandler(w, r)
}

func (m *Monitor) rootHandler(w http.ResponseWriter, r *http.Request) {
	m.jsonHandle(m.getAspectsResults(), w, r)
}

func (m *Monitor) aspectHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if a, ok := m.Aspects[name]; ok {
		m.jsonHandle(a.GetStats(), w, r)
	}
}

func (m *Monitor) jsonHandle(data interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (m *Monitor) getAspectsResults() map[string]interface{} {
	m.Lock()
	defer m.Unlock()

	r := make(map[string]interface{}, 0)
	for k, a := range m.Aspects {
		if a.InRoot() {
			r[k] = a.GetStats()
		}
	}

	return r
}
