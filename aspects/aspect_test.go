package aspects

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuntimeAspect(t *testing.T) {
	a := &RuntimeAspect{}
	r := a.GetStats()

	assert.NotEqual(t, r.(*RuntimeAspectData).CpuNum, 0)
}

func TestMemoryAspect(t *testing.T) {
	a := &MemoryAspect{}
	r := a.GetStats()

	assert.NotEqual(t, r.(*runtime.MemStats).Frees, 0)
}
