package warehouse

type Observer interface {
	Observe(subject any)
}

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify(subject any)
}
