package aspects

import (
	"runtime"
)

type MemoryAspect struct {
	ShowInRoot bool
}

func NewMemoryAspect(inRoot bool) *MemoryAspect {
	return &MemoryAspect{ShowInRoot: inRoot}
}

func (a *MemoryAspect) GetStats() interface{} {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return &mem
}

func (a *MemoryAspect) Name() string {
	return "MemStats"
}

func (a *MemoryAspect) InRoot() bool {
	return a.ShowInRoot
}
