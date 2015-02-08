package aspects

type Aspect interface {
	Get() interface{}
	Name() string
}
