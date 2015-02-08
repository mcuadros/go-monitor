package aspects

type Aspect interface {
	GetStats() interface{}
	Name() string
	InRoot() bool
}
