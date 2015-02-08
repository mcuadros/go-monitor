package monitor

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMonitor(t *testing.T) {
	m := NewMonitor(":9000")
	assert.Len(t, m.Aspects, 2)
}

func TestMonitor_getAspectsResults(t *testing.T) {
	m := NewMonitor(":9000")
	r := m.getAspectsResults()

	assert.Len(t, r, 2)
	assert.NotNil(t, r["MemStats"].(*runtime.MemStats).Alloc)
}
