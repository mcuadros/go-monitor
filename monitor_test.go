package monitor

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMonitor(t *testing.T) {
	m := NewMonitor(":9000")
	assert.Len(t, m.Aspects, 3)
}

func TestMonitor_getAllAspectsResults(t *testing.T) {
	m := NewMonitor(":9000")
	r := m.getAllAspectsResults()

	r = m.getAspectsResults("MemStats,Runtime")
	assert.Len(t, r, 2)
}

func TestMonitor_getAspectsResults(t *testing.T) {
	m := NewMonitor(":9000")
	r := m.getAspectsResults("MemStats")

	assert.Len(t, r, 1)
	assert.NotNil(t, r["MemStats"].(*runtime.MemStats).Alloc)

	r = m.getAspectsResults("MemStats,Runtime")
	assert.Len(t, r, 2)
}
