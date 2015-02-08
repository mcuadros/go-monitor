package aspects

import (
	"runtime"
)

type RuntimeAspectData struct {
	GoVersion    string
	GoOs         string
	GoArch       string
	CpuNum       int
	GoroutineNum int
	Gomaxprocs   int
	CgoCallNum   int64
}

type RuntimeAspect struct{}

func (a *RuntimeAspect) Get() interface{} {
	return &RuntimeAspectData{
		GoVersion:    runtime.Version(),
		GoOs:         runtime.GOOS,
		GoArch:       runtime.GOARCH,
		CpuNum:       runtime.NumCPU(),
		GoroutineNum: runtime.NumGoroutine(),
		Gomaxprocs:   runtime.GOMAXPROCS(0),
		CgoCallNum:   runtime.NumCgoCall(),
	}
}

func (a *RuntimeAspect) Name() string {
	return "Runtime"
}
