go-monitor [![Build Status](https://travis-ci.org/mcuadros/go-monitor.png?branch=master)](https://travis-ci.org/mcuadros/go-monitor) [![GoDoc](http://godoc.org/gopkg.in/mcuadros/go-monitor.v1?status.png)](http://godoc.org/gopkg.in/mcuadros/go-monitor.v1) [![GitHub release](https://img.shields.io/github/release/mcuadros/go-monitor.svg)](https://github.com/mcuadros/go-monitor/releases)
==============================

The main goal of `go-monitor` is provide a simple and extensible way to build monitorizable long term execution processes or daemons via HTTP.

Thanks to the defaults `aspects` you can monitorize parameters as runtime, memory, etc. for any Go processes and daemons. As well you can create your custom `aspects` for monitorize custom parameters from your applications.


Installation
------------

The recommended way to install go-monitor

```
go get gopkg.in/mcuadros/go-monitor.v1
```

Examples
--------

## Default Monitor

Import the package:

```go
import "gopkg.in/mcuadros/go-monitor.v1"
```

Start the monitor just before of the bootstrap of your code:

```go
m := monitor.NewMonitor(":9000")
m.Start()
```

Now just try `curl http://localhost:9000/`

```
{
  "MemStats": {
    "Alloc": 7716521256,
    "TotalAlloc": 1935822232552,
    "Sys": 46882078488,
    ...
  },
  "Runtime": {
    "GoVersion": "go1.3.3",
    "GoOs": "linux",
    "GoArch": "amd64",
    "CpuNum": 24,
    "GoroutineNum": 21196,
    "Gomaxprocs": 24,
    "CgoCallNum": 111584
  }
}
```

At the `/` you can find all the aspects that are loaded by default, you can request other aspects through the URL `/<aspect-name>,<aspect-name>`

## Custom Monitor

Define your custom aspect, in this case just a simple one that count the number of hits on it.

```go
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
```

Now just add the `CustomAspect` to the monitor and run it.

```go
m := monitor.NewMonitor(":9000")
m.AddAspect(&CustomAspect{})
m.Start()
```

Hit `http://localhost:9000/Custom` and obtain:
```
{
  "Custom": 5
}
```


License
-------

MIT, see [LICENSE](LICENSE)
