package monitor

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/mcuadros/go-monitor/aspects"
)

type Monitor struct {
	Addr    string
	Aspects []aspects.Aspect

	server *http.Server
	mux    *http.ServeMux

	sync.Mutex
}

func NewMonitor(addr string) *Monitor {
	return &Monitor{
		Addr: addr,
		Aspects: []aspects.Aspect{
			&aspects.MemoryAspect{},
			&aspects.RuntimeAspect{},
		},
	}
}

func (m *Monitor) Start() error {
	m.buildServer()
	return m.server.ListenAndServe()
}

func (m *Monitor) buildServer() {
	m.mux = http.NewServeMux()
	m.mux.HandleFunc("/", m.handle)

	m.server = &http.Server{Addr: m.Addr, Handler: m.mux}
}

func (m *Monitor) handle(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(m.getAspectsResults(), "  ", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (m *Monitor) getAspectsResults() map[string]interface{} {
	m.Lock()
	defer m.Unlock()

	r := make(map[string]interface{}, 0)
	for _, a := range m.Aspects {
		r[a.Name()] = a.Get()
	}

	return r
}
