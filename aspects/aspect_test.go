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

func TestTimeAspect(t *testing.T) {
	a := NewTimeAspect(true)
	r := a.GetStats()

	assert.NotEqual(t, r.(*TimeAspectData).StartTime.Nanosecond(), 0)
	assert.NotEqual(t, r.(*TimeAspectData).RequestTime.Nanosecond(), 0)
	assert.True(t, a.InRoot())
}
