go-monitor [![Build Status](https://travis-ci.org/mcuadros/go-monitor.png?branch=master)](https://travis-ci.org/mcuadros/go-monitor) [![GoDoc](http://godoc.org/github.com/mcuadros/go-monitor?status.png)](http://godoc.org/github.com/mcuadros/go-monitor)
==============================

go-monitor builds easy and extensible monitorization (runtime, memory, etc) via HTTP for Go processes and daemons.


Installation
------------

The recommended way to install go-monitor

```
go get github.com/mcuadros/go-monitor
```

Examples
--------

Import the package:

```go
import "github.com/mcuadros/go-monitor"
```

Start the monitor just before of the bootstrap of your code:

```go
m := monitor.NewMonitor(":9000")
go m.Start()
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

License
-------

MIT, see [LICENSE](LICENSE)
