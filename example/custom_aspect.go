package main

import "gopkg.in/mcuadros/go-monitor.v1"

func main() {
	m := monitor.NewMonitor(":9000")
	m.AddAspect(&CustomAspect{})
	m.Start()
}

type CustomAspect struct {
	Count int
}

func (a *CustomAspect) GetStats() interface{} {
	a.Count++
	return a.Count
}

func (a *CustomAspect) Name() string {
	return "Custom"
}

func (a *CustomAspect) InRoot() bool {
	return false
}
