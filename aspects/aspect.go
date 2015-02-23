package aspects

type Aspect interface {
	//GetStats returns any marchable object with the stats
	GetStats() interface{}
	//Name returns the name of the Aspect
	Name() string
	//InRoot returns a bool, if `true` this aspect will be used at `/`
	InRoot() bool
}
