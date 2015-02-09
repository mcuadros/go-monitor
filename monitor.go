package monitor

import (
	"encoding/json"
	"net/http"
	"strings"
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
	m.AddAspect(aspects.NewTimeAspect(true))

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
