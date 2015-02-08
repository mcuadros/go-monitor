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

type RuntimeAspect struct {
	ShowInRoot bool
}

func (a *RuntimeAspect) GetStats() interface{} {
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

func (a *RuntimeAspect) InRoot() bool {
	return a.ShowInRoot
}
