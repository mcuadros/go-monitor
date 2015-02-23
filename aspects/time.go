package aspects

import (
	"time"
)

type TimeAspectData struct {
	StartTime   time.Time
	RequestTime time.Time
}

type TimeAspect struct {
	StartTime  time.Time
	ShowInRoot bool
}

//NewTimeAspect returns the current time and the running time
func NewTimeAspect(inRoot bool) *TimeAspect {
	return &TimeAspect{StartTime: time.Now(), ShowInRoot: inRoot}
}

func (a *TimeAspect) GetStats() interface{} {
	return &TimeAspectData{
		StartTime:   a.StartTime,
		RequestTime: time.Now(),
	}
}

func (a *TimeAspect) Name() string {
	return "Time"
}

func (a *TimeAspect) InRoot() bool {
	return a.ShowInRoot
}
