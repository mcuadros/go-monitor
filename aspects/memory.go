package aspects

import (
	"runtime"
)

type Foo struct {
	runtime.MemStats
}

type MemoryAspect struct{}

func (a *MemoryAspect) Get() interface{} {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return &mem
}

func (a *MemoryAspect) Name() string {
	return "MemStats"
}
